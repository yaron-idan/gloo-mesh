package tests

import (
	"context"
	"log"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	networkingv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/test/data"
	. "github.com/solo-io/gloo-mesh/test/e2e"
	"github.com/solo-io/gloo-mesh/test/utils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	err                 error
	VirtualMesh         *networkingv1alpha2.VirtualMesh
	VirtualMeshManifest utils.Manifest
)

// Before running tests, federate the two clusters by creating a VirtualMesh with mTLS enabled.
func SetupClustersAndFederation(customDeployFuc func()) {
	VirtualMeshManifest, err = utils.NewManifest("virtualmesh.yaml")
	Expect(err).NotTo(HaveOccurred())

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	if customDeployFuc != nil {
		os.Setenv("SKIP_DEPLOY_FROM_SOURCE", "1")
	}
	/* env := */ StartEnvOnce(ctx)

	if customDeployFuc != nil {
		customDeployFuc()
	}

	var err error
	dynamicClient, err = client.New(GetEnv().Management.Config, client.Options{})
	Expect(err).NotTo(HaveOccurred())

	federateClusters(dynamicClient)
}

func federateClusters(dynamicClient client.Client) {
	VirtualMesh = data.SelfSignedVirtualMesh(
		"bookinfo-federation",
		BookinfoNamespace,
		[]*v1.ObjectRef{
			masterMesh,
			remoteMesh,
		})

	err = VirtualMeshManifest.AppendResources(VirtualMesh)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() error {
		// retry for the case where we're running against rbac-webhook
		return VirtualMeshManifest.KubeApply(BookinfoNamespace)
	}, "60s", "1s").ShouldNot(HaveOccurred())

	// ensure status is updated
	utils.AssertVirtualMeshStatuses(dynamicClient, BookinfoNamespace)

	// check we can hit the remote service
	// give 5 minutes because the workflow depends on restarting pods
	// which can take several minutes
	Eventually(curlRemoteReviews, "5m", "2s").Should(ContainSubstring("200 OK"))
}

func TeardownFederationAndClusters() {
	err = VirtualMeshManifest.KubeDelete(BookinfoNamespace)
	if err != nil {
		// this is expected to fail in gloo-mesh-enterprise-helm tests as they run the rbac webhook which disables ability to delete this manifest
		log.Printf("warn: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if os.Getenv("NO_CLEANUP") != "" {
		return
	}
	_ = ClearEnv(ctx)
}

// initialize all tests in suite
// should be called from init() function or top level var
func InitializeTests() bool {
	var (
		_ = Describe("AccessPolicy", AccessPolicyTest)
		_ = Describe("FailoverService", FailoverServiceTest)
		_ = Describe("Federation", FederationTest)
		_ = Describe("Networking Extensions", NetworkingExtensionsTest)
		_ = Describe("TrafficPolicy", TrafficPolicyTest)
	)
	return true
}
