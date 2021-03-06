
---
title: "role.proto"
---

## Package : `rbac.enterprise.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for role.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## role.proto


## Table of Contents
  - [RoleBindingSpec](#rbac.enterprise.mesh.gloo.solo.io.RoleBindingSpec)
  - [RoleBindingStatus](#rbac.enterprise.mesh.gloo.solo.io.RoleBindingStatus)
  - [RoleSpec](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec)
  - [RoleSpec.AccessPolicyScope](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.AccessPolicyScope)
  - [RoleSpec.FailoverServiceScope](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.FailoverServiceScope)
  - [RoleSpec.TrafficPolicyScope](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope)
  - [RoleSpec.VirtualMeshScope](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope)
  - [RoleSpec.WasmDeploymentScope](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.WasmDeploymentScope)
  - [RoleStatus](#rbac.enterprise.mesh.gloo.solo.io.RoleStatus)

  - [RoleSpec.TrafficPolicyScope.TrafficPolicyActions](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope.TrafficPolicyActions)
  - [RoleSpec.VirtualMeshScope.VirtualMeshActions](#rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope.VirtualMeshActions)






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleBindingSpec"></a>

### RoleBindingSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subjects | []core.skv2.solo.io.TypedObjectRef | repeated | reference to users or groups to apply the Gloo Mesh Role to |
| roleRef | core.skv2.solo.io.ObjectRef |  | reference to a Gloo Mesh Role |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleBindingStatus"></a>

### RoleBindingStatus







<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec"></a>

### RoleSpec
A role represents a set of permissions for creating, updating, and deleting Gloo Mesh configuration. A role consists of a set of scopes for each policy type. The permission granularity is defined at the field level for TrafficPolicy and VirtualMesh and at the object level for AccessPolicy and FailoverService.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trafficPolicyScopes | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope | repeated | A set of TrafficPolicy configuration permissions. Permission granularity is defined at the field level. |
| virtualMeshScopes | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope | repeated | A set of VirtualMesh configuration permissions. Permission granularity is defined at the field level. |
| accessPolicyScopes | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.AccessPolicyScope | repeated | A set of AccessPolicy configuration permissions. Permission granularity is defined at the object level. |
| failoverServiceScopes | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.FailoverServiceScope | repeated | A set of FailoverService configuration permissions. Permission granularity is defined at the object level. |
| wasmDeploymentScopes | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.WasmDeploymentScope | repeated | A set of WasmDeployment configuration permissions. Permission granularity is defined at the object level. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.AccessPolicyScope"></a>

### RoleSpec.AccessPolicyScope
Represents permissions for configuring AccessPolicies.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identitySelectors | []networking.mesh.gloo.solo.io.IdentitySelector | repeated | A list of permitted identity selectors. |
| trafficTargetSelectors | []networking.mesh.gloo.solo.io.TrafficTargetSelector | repeated | A list of permitted traffic target selectors. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.FailoverServiceScope"></a>

### RoleSpec.FailoverServiceScope
Represents permissions for configuring FailoverServices.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meshRefs | []core.skv2.solo.io.ObjectRef | repeated | A list of permitted mesh references. |
| backingServices | []networking.mesh.gloo.solo.io.FailoverServiceSpec.BackingService | repeated | A list of permitted backing services. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope"></a>

### RoleSpec.TrafficPolicyScope
Represents permissions for configuring TrafficPolicies.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trafficPolicyActions | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope.TrafficPolicyActions | repeated | A list of permitted TrafficPolicy configuration actions. |
| trafficTargetSelectors | []networking.mesh.gloo.solo.io.TrafficTargetSelector | repeated | A list of permitted traffic target selectors. |
| workloadSelectors | []networking.mesh.gloo.solo.io.WorkloadSelector | repeated | A list of permitted workload selectors. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope"></a>

### RoleSpec.VirtualMeshScope
Represents permissions for configuring VirtualMeshes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtualMeshActions | []rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope.VirtualMeshActions | repeated | A list of permitted VirtualMesh configuration actions. |
| meshRefs | []core.skv2.solo.io.ObjectRef | repeated | A list of permitted mesh references. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.WasmDeploymentScope"></a>

### RoleSpec.WasmDeploymentScope
Represents permissions for configuring WasmDeployments.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| workloadSelectors | []networking.mesh.gloo.solo.io.WorkloadSelector | repeated | A list of permitted workload selectors. |






<a name="rbac.enterprise.mesh.gloo.solo.io.RoleStatus"></a>

### RoleStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  |  |





 <!-- end messages -->


<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.TrafficPolicyScope.TrafficPolicyActions"></a>

### RoleSpec.TrafficPolicyScope.TrafficPolicyActions
Enums representing fields on the TrafficPolicy CRD.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_TRAFFIC_POLICY_ACTION | 0 |  |
| ALL | 1 |  |
| TRAFFIC_SHIFT | 2 |  |
| FAULT_INJECTION | 3 |  |
| REQUEST_TIMEOUT | 4 |  |
| RETRIES | 5 |  |
| CORS_POLICY | 6 |  |
| MIRROR | 7 |  |
| HEADER_MANIPULATION | 8 |  |
| OUTLIER_DETECTION | 9 |  |
| MTLS_CONFIG | 10 |  |



<a name="rbac.enterprise.mesh.gloo.solo.io.RoleSpec.VirtualMeshScope.VirtualMeshActions"></a>

### RoleSpec.VirtualMeshScope.VirtualMeshActions
Enums representing fields on the VirtualMesh CRD.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_VIRTUAL_MESH_ACTION | 0 |  |
| ALL | 1 |  |
| MTLS_CONFIG | 2 |  |
| FEDERATION | 3 |  |
| GLOBAL_ACCESS_POLICY | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

