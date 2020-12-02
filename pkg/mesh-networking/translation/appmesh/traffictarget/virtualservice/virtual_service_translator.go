package virtualservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	appmeshv1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	"github.com/solo-io/go-utils/contextutils"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/utils/workloadutils"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Translator interface {
	Translate(
		ctx context.Context,
		in input.Snapshot,
		trafficTarget *discoveryv1alpha2.TrafficTarget,
		virtualRouter *appmeshv1beta2.VirtualRouter,
		reporter reporting.Reporter,
	) []*appmeshv1beta2.VirtualService
}

type translator struct{}

func NewVirtualServiceTranslator() Translator {
	return &translator{}
}

func (t *translator) Translate(
	ctx context.Context,
	in input.Snapshot,
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	virtualRouter *appmeshv1beta2.VirtualRouter,
	reporter reporting.Reporter,
) []*appmeshv1beta2.VirtualService {

	kubeService := trafficTarget.Spec.GetKubeService()
	if kubeService == nil {
		// TODO non kube services currently unsupported
		return nil
	}

	backingWorkloads := workloadutils.FindBackingWorkloads(trafficTarget.Spec.GetKubeService(), in.Workloads())
	if len(backingWorkloads) == 0 {
		contextutils.LoggerFrom(ctx).Warnf("Found no backing workloads for traffic target %s", sets.Key(trafficTarget))
		return nil
	}

	// Create a virtual service for all backing workloads
	virtualServices := make([]*appmeshv1beta2.VirtualService, 0, len(backingWorkloads))
	for _, workload := range backingWorkloads {
		arn := workload.Spec.AppMesh.VirtualNodeArn
		meta := metautils.TranslatedObjectMeta(
			mergeRefs(trafficTarget.Spec.GetKubeService().Ref, workload),
			trafficTarget.Annotations,
		)

		virtualServices = append(virtualServices, getVirtualService(meta, virtualRouter, arn))
	}

	return virtualServices
}

func getVirtualService(
	meta metav1.ObjectMeta,
	virtualRouter *appmeshv1beta2.VirtualRouter,
	arn string,
) *appmeshv1beta2.VirtualService {
	var provider *appmeshv1beta2.VirtualServiceProvider
	if virtualRouter != nil {
		provider = &appmeshv1beta2.VirtualServiceProvider{
			VirtualRouter: &appmeshv1beta2.VirtualRouterServiceProvider{
				VirtualRouterRef: &appmeshv1beta2.VirtualRouterReference{
					Namespace: &meta.Namespace,
					Name:      meta.Name,
				},
			},
		}
	} else {
		provider = &appmeshv1beta2.VirtualServiceProvider{
			VirtualNode: &appmeshv1beta2.VirtualNodeServiceProvider{
				VirtualNodeARN: &arn,
			},
		}
	}

	// This is the default name written back by the AWS controller.
	// We must provide it explicitly, else the App Mesh controller's
	// validating admission webhook will reject our changes on update.
	awsName := fmt.Sprintf("%s.%s", meta.Name, meta.Namespace)
	return &appmeshv1beta2.VirtualService{
		ObjectMeta: meta,
		Spec: v1beta2.VirtualServiceSpec{
			AWSName:  &awsName,
			Provider: provider,
		},
	}
}

func mergeRefs(refs ...ezkube.ClusterResourceId) ezkube.ClusterResourceId {
	output := &v1.ClusterObjectRef{}

	for _, ref := range refs {
		output = &v1.ClusterObjectRef{
			Name:        strings.Join([]string{output.Name, ref.GetName()}, "-"),
			Namespace:   strings.Join([]string{output.Namespace, ref.GetNamespace()}, "-"),
			ClusterName: strings.Join([]string{output.ClusterName, ref.GetClusterName()}, "-"),
		}
	}
}
