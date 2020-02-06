package linkerd

import (
	"context"
	"strings"

	"github.com/google/wire"
	"github.com/rotisserie/eris"
	mp_v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1"
	v1alpha1_types "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/mesh-projects/pkg/common/docker"
	"github.com/solo-io/mesh-projects/pkg/env"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/discovery"
	k8s_apps_v1 "k8s.io/api/apps/v1"
	k8s_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	WireProviderSet = wire.NewSet(
		NewLinkerdMeshFinder,
	)
	DiscoveryLabels = map[string]string{
		"discovered_by": "linkerd-mesh-discovery",
	}
	UnexpectedControllerImageName = func(err error, imageName string) error {
		return eris.Wrapf(err, "invalid or unexpected image name format for linkerd controller: %s", imageName)
	}
)

// disambiguates this MeshFinder from the other MeshFinder implementations so that wire stays happy
type LinkerdMeshFinder discovery.MeshFinder

func NewLinkerdMeshFinder(imageNameParser docker.ImageNameParser) LinkerdMeshFinder {
	return &linkerdMeshFinder{
		imageNameParser: imageNameParser,
	}
}

type linkerdMeshFinder struct {
	imageNameParser docker.ImageNameParser
}

func (l *linkerdMeshFinder) ScanDeployment(_ context.Context, deployment *k8s_apps_v1.Deployment) (*mp_v1alpha1.Mesh, error) {
	linkerdController, err := l.detectLinkerdController(deployment)

	if err != nil {
		return nil, err
	}

	if linkerdController == nil {
		return nil, nil
	}

	return &mp_v1alpha1.Mesh{
		ObjectMeta: k8s_meta_v1.ObjectMeta{
			Name:      linkerdController.name(),
			Namespace: env.DefaultWriteNamespace,
			Labels:    DiscoveryLabels,
		},
		Spec: v1alpha1_types.MeshSpec{
			MeshType: &v1alpha1_types.MeshSpec_Linkerd{
				Linkerd: &v1alpha1_types.LinkerdMesh{
					Installation: &v1alpha1_types.MeshInstallation{
						InstallationNamespace: deployment.GetNamespace(),
						Version:               linkerdController.version,
					},
				},
			},
			Cluster: &v1alpha1_types.ResourceRef{
				Name:      deployment.GetClusterName(),
				Namespace: env.DefaultWriteNamespace,
			},
		},
	}, nil
}

func (l *linkerdMeshFinder) detectLinkerdController(deployment *k8s_apps_v1.Deployment) (*linkerdControllerDeployment, error) {
	var linkerdController *linkerdControllerDeployment

	for _, container := range deployment.Spec.Template.Spec.Containers {
		if strings.Contains(container.Image, "linkerd-io/controller") {
			// TODO there can be > 1 controller image per pod, do we care?
			parsedImage, err := l.imageNameParser.Parse(container.Image)
			if err != nil {
				return nil, UnexpectedControllerImageName(err, container.Image)
			}

			version := parsedImage.Tag
			if parsedImage.Digest != "" {
				version = parsedImage.Digest
			}
			linkerdController = &linkerdControllerDeployment{version: version, namespace: deployment.Namespace, cluster: deployment.ClusterName}
		}
	}

	return linkerdController, nil
}

type linkerdControllerDeployment struct {
	version, namespace, cluster string
}

func (c linkerdControllerDeployment) name() string {
	if c.cluster == "" {
		return "linkerd-" + c.namespace
	}
	// TODO cluster is not restricted to kube name scheme, kebab it
	return "linkerd-" + c.namespace + "-" + c.cluster
}