// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	settings_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	settings_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"

	discovery_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	discovery_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"

	networking_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	networking_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	appmesh_k8s_aws_v1beta2_sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

type InputSnapshotManualBuilder struct {
	name string

	settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet

	trafficTargets discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet
	workloads      discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet
	meshes         discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet

	trafficPolicies  networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet
	accessPolicies   networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet
	virtualMeshes    networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet
	failoverServices networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet

	virtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet
	virtualNodes    appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet

	secrets v1_sets.SecretSet

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		settings: settings_mesh_gloo_solo_io_v1alpha2_sets.NewSettingsSet(),

		trafficTargets: discovery_mesh_gloo_solo_io_v1alpha2_sets.NewTrafficTargetSet(),
		workloads:      discovery_mesh_gloo_solo_io_v1alpha2_sets.NewWorkloadSet(),
		meshes:         discovery_mesh_gloo_solo_io_v1alpha2_sets.NewMeshSet(),

		trafficPolicies:  networking_mesh_gloo_solo_io_v1alpha2_sets.NewTrafficPolicySet(),
		accessPolicies:   networking_mesh_gloo_solo_io_v1alpha2_sets.NewAccessPolicySet(),
		virtualMeshes:    networking_mesh_gloo_solo_io_v1alpha2_sets.NewVirtualMeshSet(),
		failoverServices: networking_mesh_gloo_solo_io_v1alpha2_sets.NewFailoverServiceSet(),

		virtualServices: appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet(),
		virtualNodes:    appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet(),

		secrets: v1_sets.NewSecretSet(),

		kubernetesClusters: multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.settings,

		i.trafficTargets,
		i.workloads,
		i.meshes,

		i.trafficPolicies,
		i.accessPolicies,
		i.virtualMeshes,
		i.failoverServices,

		i.virtualServices,
		i.virtualNodes,

		i.secrets,

		i.kubernetesClusters,
	)
}
func (i *InputSnapshotManualBuilder) AddSettings(settings []*settings_mesh_gloo_solo_io_v1alpha2.Settings) *InputSnapshotManualBuilder {
	i.settings.Insert(settings...)
	return i
}
func (i *InputSnapshotManualBuilder) AddTrafficTargets(trafficTargets []*discovery_mesh_gloo_solo_io_v1alpha2.TrafficTarget) *InputSnapshotManualBuilder {
	i.trafficTargets.Insert(trafficTargets...)
	return i
}
func (i *InputSnapshotManualBuilder) AddWorkloads(workloads []*discovery_mesh_gloo_solo_io_v1alpha2.Workload) *InputSnapshotManualBuilder {
	i.workloads.Insert(workloads...)
	return i
}
func (i *InputSnapshotManualBuilder) AddMeshes(meshes []*discovery_mesh_gloo_solo_io_v1alpha2.Mesh) *InputSnapshotManualBuilder {
	i.meshes.Insert(meshes...)
	return i
}
func (i *InputSnapshotManualBuilder) AddTrafficPolicies(trafficPolicies []*networking_mesh_gloo_solo_io_v1alpha2.TrafficPolicy) *InputSnapshotManualBuilder {
	i.trafficPolicies.Insert(trafficPolicies...)
	return i
}
func (i *InputSnapshotManualBuilder) AddAccessPolicies(accessPolicies []*networking_mesh_gloo_solo_io_v1alpha2.AccessPolicy) *InputSnapshotManualBuilder {
	i.accessPolicies.Insert(accessPolicies...)
	return i
}
func (i *InputSnapshotManualBuilder) AddVirtualMeshes(virtualMeshes []*networking_mesh_gloo_solo_io_v1alpha2.VirtualMesh) *InputSnapshotManualBuilder {
	i.virtualMeshes.Insert(virtualMeshes...)
	return i
}
func (i *InputSnapshotManualBuilder) AddFailoverServices(failoverServices []*networking_mesh_gloo_solo_io_v1alpha2.FailoverService) *InputSnapshotManualBuilder {
	i.failoverServices.Insert(failoverServices...)
	return i
}
func (i *InputSnapshotManualBuilder) AddVirtualServices(virtualServices []*appmesh_k8s_aws_v1beta2.VirtualService) *InputSnapshotManualBuilder {
	i.virtualServices.Insert(virtualServices...)
	return i
}
func (i *InputSnapshotManualBuilder) AddVirtualNodes(virtualNodes []*appmesh_k8s_aws_v1beta2.VirtualNode) *InputSnapshotManualBuilder {
	i.virtualNodes.Insert(virtualNodes...)
	return i
}
func (i *InputSnapshotManualBuilder) AddSecrets(secrets []*v1.Secret) *InputSnapshotManualBuilder {
	i.secrets.Insert(secrets...)
	return i
}
func (i *InputSnapshotManualBuilder) AddKubernetesClusters(kubernetesClusters []*multicluster_solo_io_v1alpha1.KubernetesCluster) *InputSnapshotManualBuilder {
	i.kubernetesClusters.Insert(kubernetesClusters...)
	return i
}