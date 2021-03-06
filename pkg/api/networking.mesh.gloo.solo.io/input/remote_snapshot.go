// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./remote_snapshot.go -destination mocks/remote_snapshot.go

// The Input RemoteSnapshot contains the set of all:
// * IssuedCertificates
// * PodBounceDirectives
// * XdsConfigs
// * DestinationRules
// * EnvoyFilters
// * Gateways
// * ServiceEntries
// * VirtualServices
// * AuthorizationPolicies
// * ConfigMaps
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the RemoteSnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/sets"

	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1"
	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1/sets"

	networking_istio_io_v1alpha3 "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3"
	networking_istio_io_v1alpha3_sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"

	security_istio_io_v1beta1 "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1"
	security_istio_io_v1beta1_sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
)

// the snapshot of input resources consumed by translation
type RemoteSnapshot interface {

	// return the set of input IssuedCertificates
	IssuedCertificates() certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	// return the set of input PodBounceDirectives
	PodBounceDirectives() certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet

	// return the set of input XdsConfigs
	XdsConfigs() xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.XdsConfigSet

	// return the set of input DestinationRules
	DestinationRules() networking_istio_io_v1alpha3_sets.DestinationRuleSet
	// return the set of input EnvoyFilters
	EnvoyFilters() networking_istio_io_v1alpha3_sets.EnvoyFilterSet
	// return the set of input Gateways
	Gateways() networking_istio_io_v1alpha3_sets.GatewaySet
	// return the set of input ServiceEntries
	ServiceEntries() networking_istio_io_v1alpha3_sets.ServiceEntrySet
	// return the set of input VirtualServices
	VirtualServices() networking_istio_io_v1alpha3_sets.VirtualServiceSet

	// return the set of input AuthorizationPolicies
	AuthorizationPolicies() security_istio_io_v1beta1_sets.AuthorizationPolicySet

	// return the set of input ConfigMaps
	ConfigMaps() v1_sets.ConfigMapSet
	// update the status of all input objects which support
	// the Status subresource (across multiple clusters)
	SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts RemoteSyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type RemoteSyncStatusOptions struct {

	// sync status of IssuedCertificate objects
	IssuedCertificate bool
	// sync status of PodBounceDirective objects
	PodBounceDirective bool

	// sync status of XdsConfig objects
	XdsConfig bool

	// sync status of DestinationRule objects
	DestinationRule bool
	// sync status of EnvoyFilter objects
	EnvoyFilter bool
	// sync status of Gateway objects
	Gateway bool
	// sync status of ServiceEntry objects
	ServiceEntry bool
	// sync status of VirtualService objects
	VirtualService bool

	// sync status of AuthorizationPolicy objects
	AuthorizationPolicy bool

	// sync status of ConfigMap objects
	ConfigMap bool
}

type snapshotRemote struct {
	name string

	issuedCertificates  certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	podBounceDirectives certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet

	xdsConfigs xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.XdsConfigSet

	destinationRules networking_istio_io_v1alpha3_sets.DestinationRuleSet
	envoyFilters     networking_istio_io_v1alpha3_sets.EnvoyFilterSet
	gateways         networking_istio_io_v1alpha3_sets.GatewaySet
	serviceEntries   networking_istio_io_v1alpha3_sets.ServiceEntrySet
	virtualServices  networking_istio_io_v1alpha3_sets.VirtualServiceSet

	authorizationPolicies security_istio_io_v1beta1_sets.AuthorizationPolicySet

	configMaps v1_sets.ConfigMapSet
}

func NewRemoteSnapshot(
	name string,

	issuedCertificates certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet,
	podBounceDirectives certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet,

	xdsConfigs xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.XdsConfigSet,

	destinationRules networking_istio_io_v1alpha3_sets.DestinationRuleSet,
	envoyFilters networking_istio_io_v1alpha3_sets.EnvoyFilterSet,
	gateways networking_istio_io_v1alpha3_sets.GatewaySet,
	serviceEntries networking_istio_io_v1alpha3_sets.ServiceEntrySet,
	virtualServices networking_istio_io_v1alpha3_sets.VirtualServiceSet,

	authorizationPolicies security_istio_io_v1beta1_sets.AuthorizationPolicySet,

	configMaps v1_sets.ConfigMapSet,

) RemoteSnapshot {
	return &snapshotRemote{
		name: name,

		issuedCertificates:    issuedCertificates,
		podBounceDirectives:   podBounceDirectives,
		xdsConfigs:            xdsConfigs,
		destinationRules:      destinationRules,
		envoyFilters:          envoyFilters,
		gateways:              gateways,
		serviceEntries:        serviceEntries,
		virtualServices:       virtualServices,
		authorizationPolicies: authorizationPolicies,
		configMaps:            configMaps,
	}
}

func (s snapshotRemote) IssuedCertificates() certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet {
	return s.issuedCertificates
}

func (s snapshotRemote) PodBounceDirectives() certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet {
	return s.podBounceDirectives
}

func (s snapshotRemote) XdsConfigs() xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.XdsConfigSet {
	return s.xdsConfigs
}

func (s snapshotRemote) DestinationRules() networking_istio_io_v1alpha3_sets.DestinationRuleSet {
	return s.destinationRules
}

func (s snapshotRemote) EnvoyFilters() networking_istio_io_v1alpha3_sets.EnvoyFilterSet {
	return s.envoyFilters
}

func (s snapshotRemote) Gateways() networking_istio_io_v1alpha3_sets.GatewaySet {
	return s.gateways
}

func (s snapshotRemote) ServiceEntries() networking_istio_io_v1alpha3_sets.ServiceEntrySet {
	return s.serviceEntries
}

func (s snapshotRemote) VirtualServices() networking_istio_io_v1alpha3_sets.VirtualServiceSet {
	return s.virtualServices
}

func (s snapshotRemote) AuthorizationPolicies() security_istio_io_v1beta1_sets.AuthorizationPolicySet {
	return s.authorizationPolicies
}

func (s snapshotRemote) ConfigMaps() v1_sets.ConfigMapSet {
	return s.configMaps
}

func (s snapshotRemote) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts RemoteSyncStatusOptions) error {
	var errs error

	if opts.IssuedCertificate {
		for _, obj := range s.IssuedCertificates().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.PodBounceDirective {
		for _, obj := range s.PodBounceDirectives().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.XdsConfig {
		for _, obj := range s.XdsConfigs().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	return errs
}

func (s snapshotRemote) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["issuedCertificates"] = s.issuedCertificates.List()
	snapshotMap["podBounceDirectives"] = s.podBounceDirectives.List()
	snapshotMap["xdsConfigs"] = s.xdsConfigs.List()
	snapshotMap["destinationRules"] = s.destinationRules.List()
	snapshotMap["envoyFilters"] = s.envoyFilters.List()
	snapshotMap["gateways"] = s.gateways.List()
	snapshotMap["serviceEntries"] = s.serviceEntries.List()
	snapshotMap["virtualServices"] = s.virtualServices.List()
	snapshotMap["authorizationPolicies"] = s.authorizationPolicies.List()
	snapshotMap["configMaps"] = s.configMaps.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type RemoteBuilder interface {
	BuildSnapshot(ctx context.Context, name string, opts RemoteBuildOptions) (RemoteSnapshot, error)
}

// Options for building a snapshot
type RemoteBuildOptions struct {

	// List options for composing a snapshot from IssuedCertificates
	IssuedCertificates ResourceRemoteBuildOptions
	// List options for composing a snapshot from PodBounceDirectives
	PodBounceDirectives ResourceRemoteBuildOptions

	// List options for composing a snapshot from XdsConfigs
	XdsConfigs ResourceRemoteBuildOptions

	// List options for composing a snapshot from DestinationRules
	DestinationRules ResourceRemoteBuildOptions
	// List options for composing a snapshot from EnvoyFilters
	EnvoyFilters ResourceRemoteBuildOptions
	// List options for composing a snapshot from Gateways
	Gateways ResourceRemoteBuildOptions
	// List options for composing a snapshot from ServiceEntries
	ServiceEntries ResourceRemoteBuildOptions
	// List options for composing a snapshot from VirtualServices
	VirtualServices ResourceRemoteBuildOptions

	// List options for composing a snapshot from AuthorizationPolicies
	AuthorizationPolicies ResourceRemoteBuildOptions

	// List options for composing a snapshot from ConfigMaps
	ConfigMaps ResourceRemoteBuildOptions
}

// Options for reading resources of a given type
type ResourceRemoteBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterRemoteBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterRemoteBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) RemoteBuilder {
	return &multiClusterRemoteBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterRemoteBuilder) BuildSnapshot(ctx context.Context, name string, opts RemoteBuildOptions) (RemoteSnapshot, error) {

	issuedCertificates := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewIssuedCertificateSet()
	podBounceDirectives := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewPodBounceDirectiveSet()

	xdsConfigs := xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.NewXdsConfigSet()

	destinationRules := networking_istio_io_v1alpha3_sets.NewDestinationRuleSet()
	envoyFilters := networking_istio_io_v1alpha3_sets.NewEnvoyFilterSet()
	gateways := networking_istio_io_v1alpha3_sets.NewGatewaySet()
	serviceEntries := networking_istio_io_v1alpha3_sets.NewServiceEntrySet()
	virtualServices := networking_istio_io_v1alpha3_sets.NewVirtualServiceSet()

	authorizationPolicies := security_istio_io_v1beta1_sets.NewAuthorizationPolicySet()

	configMaps := v1_sets.NewConfigMapSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertIssuedCertificatesFromCluster(ctx, cluster, issuedCertificates, opts.IssuedCertificates); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertPodBounceDirectivesFromCluster(ctx, cluster, podBounceDirectives, opts.PodBounceDirectives); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertXdsConfigsFromCluster(ctx, cluster, xdsConfigs, opts.XdsConfigs); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDestinationRulesFromCluster(ctx, cluster, destinationRules, opts.DestinationRules); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertEnvoyFiltersFromCluster(ctx, cluster, envoyFilters, opts.EnvoyFilters); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertGatewaysFromCluster(ctx, cluster, gateways, opts.Gateways); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertServiceEntriesFromCluster(ctx, cluster, serviceEntries, opts.ServiceEntries); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertVirtualServicesFromCluster(ctx, cluster, virtualServices, opts.VirtualServices); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertAuthorizationPoliciesFromCluster(ctx, cluster, authorizationPolicies, opts.AuthorizationPolicies); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertConfigMapsFromCluster(ctx, cluster, configMaps, opts.ConfigMaps); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewRemoteSnapshot(
		name,

		issuedCertificates,
		podBounceDirectives,
		xdsConfigs,
		destinationRules,
		envoyFilters,
		gateways,
		serviceEntries,
		virtualServices,
		authorizationPolicies,
		configMaps,
	)

	return outputSnap, errs
}

func (b *multiClusterRemoteBuilder) insertIssuedCertificatesFromCluster(ctx context.Context, cluster string, issuedCertificates certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet, opts ResourceRemoteBuildOptions) error {
	issuedCertificateClient, err := certificates_mesh_gloo_solo_io_v1alpha2.NewMulticlusterIssuedCertificateClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "IssuedCertificate",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	issuedCertificateList, err := issuedCertificateClient.ListIssuedCertificate(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range issuedCertificateList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		issuedCertificates.Insert(&item)
	}

	return nil
}
func (b *multiClusterRemoteBuilder) insertPodBounceDirectivesFromCluster(ctx context.Context, cluster string, podBounceDirectives certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet, opts ResourceRemoteBuildOptions) error {
	podBounceDirectiveClient, err := certificates_mesh_gloo_solo_io_v1alpha2.NewMulticlusterPodBounceDirectiveClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "PodBounceDirective",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	podBounceDirectiveList, err := podBounceDirectiveClient.ListPodBounceDirective(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range podBounceDirectiveList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		podBounceDirectives.Insert(&item)
	}

	return nil
}

func (b *multiClusterRemoteBuilder) insertXdsConfigsFromCluster(ctx context.Context, cluster string, xdsConfigs xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_sets.XdsConfigSet, opts ResourceRemoteBuildOptions) error {
	xdsConfigClient, err := xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.NewMulticlusterXdsConfigClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "xds.agent.enterprise.mesh.gloo.solo.io",
			Version: "v1alpha1",
			Kind:    "XdsConfig",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	xdsConfigList, err := xdsConfigClient.ListXdsConfig(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range xdsConfigList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		xdsConfigs.Insert(&item)
	}

	return nil
}

func (b *multiClusterRemoteBuilder) insertDestinationRulesFromCluster(ctx context.Context, cluster string, destinationRules networking_istio_io_v1alpha3_sets.DestinationRuleSet, opts ResourceRemoteBuildOptions) error {
	destinationRuleClient, err := networking_istio_io_v1alpha3.NewMulticlusterDestinationRuleClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Version: "v1alpha3",
			Kind:    "DestinationRule",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	destinationRuleList, err := destinationRuleClient.ListDestinationRule(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range destinationRuleList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		destinationRules.Insert(&item)
	}

	return nil
}
func (b *multiClusterRemoteBuilder) insertEnvoyFiltersFromCluster(ctx context.Context, cluster string, envoyFilters networking_istio_io_v1alpha3_sets.EnvoyFilterSet, opts ResourceRemoteBuildOptions) error {
	envoyFilterClient, err := networking_istio_io_v1alpha3.NewMulticlusterEnvoyFilterClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Version: "v1alpha3",
			Kind:    "EnvoyFilter",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	envoyFilterList, err := envoyFilterClient.ListEnvoyFilter(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range envoyFilterList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		envoyFilters.Insert(&item)
	}

	return nil
}
func (b *multiClusterRemoteBuilder) insertGatewaysFromCluster(ctx context.Context, cluster string, gateways networking_istio_io_v1alpha3_sets.GatewaySet, opts ResourceRemoteBuildOptions) error {
	gatewayClient, err := networking_istio_io_v1alpha3.NewMulticlusterGatewayClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Version: "v1alpha3",
			Kind:    "Gateway",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	gatewayList, err := gatewayClient.ListGateway(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range gatewayList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		gateways.Insert(&item)
	}

	return nil
}
func (b *multiClusterRemoteBuilder) insertServiceEntriesFromCluster(ctx context.Context, cluster string, serviceEntries networking_istio_io_v1alpha3_sets.ServiceEntrySet, opts ResourceRemoteBuildOptions) error {
	serviceEntryClient, err := networking_istio_io_v1alpha3.NewMulticlusterServiceEntryClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Version: "v1alpha3",
			Kind:    "ServiceEntry",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	serviceEntryList, err := serviceEntryClient.ListServiceEntry(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range serviceEntryList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		serviceEntries.Insert(&item)
	}

	return nil
}
func (b *multiClusterRemoteBuilder) insertVirtualServicesFromCluster(ctx context.Context, cluster string, virtualServices networking_istio_io_v1alpha3_sets.VirtualServiceSet, opts ResourceRemoteBuildOptions) error {
	virtualServiceClient, err := networking_istio_io_v1alpha3.NewMulticlusterVirtualServiceClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Version: "v1alpha3",
			Kind:    "VirtualService",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	virtualServiceList, err := virtualServiceClient.ListVirtualService(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range virtualServiceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		virtualServices.Insert(&item)
	}

	return nil
}

func (b *multiClusterRemoteBuilder) insertAuthorizationPoliciesFromCluster(ctx context.Context, cluster string, authorizationPolicies security_istio_io_v1beta1_sets.AuthorizationPolicySet, opts ResourceRemoteBuildOptions) error {
	authorizationPolicyClient, err := security_istio_io_v1beta1.NewMulticlusterAuthorizationPolicyClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "security.istio.io",
			Version: "v1beta1",
			Kind:    "AuthorizationPolicy",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	authorizationPolicyList, err := authorizationPolicyClient.ListAuthorizationPolicy(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range authorizationPolicyList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		authorizationPolicies.Insert(&item)
	}

	return nil
}

func (b *multiClusterRemoteBuilder) insertConfigMapsFromCluster(ctx context.Context, cluster string, configMaps v1_sets.ConfigMapSet, opts ResourceRemoteBuildOptions) error {
	configMapClient, err := v1.NewMulticlusterConfigMapClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "ConfigMap",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	configMapList, err := configMapClient.ListConfigMap(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range configMapList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		configMaps.Insert(&item)
	}

	return nil
}
