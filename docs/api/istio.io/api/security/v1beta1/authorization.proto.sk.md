
---
title: "authorization.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `istio.security.v1beta1` 
#### Types:


- [AuthorizationPolicy](#authorizationpolicy)
- [Rule](#rule)
- [From](#from)
- [To](#to)
- [Source](#source)
- [Operation](#operation)
- [Condition](#condition)
  



##### Source File: `istio.io/api/security/v1beta1/authorization.proto`





---
### AuthorizationPolicy

 
AuthorizationPolicy enables access control on workloads.

For example, the following authorization policy denies all requests to workloads
in namespace foo.

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
 name: deny-all
 namespace: foo
spec:
```

The following authorization policy allows all requests to workloads in namespace
foo.

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
 name: allow-all
 namespace: foo
spec:
 rules:
 - {}
```

<!-- go code generation tags
+kubetype-gen
+kubetype-gen:groupVersion=security.istio.io/v1beta1
+genclient
+k8s:deepcopy-gen=true
-->

```yaml
"selector": .istio.type.v1beta1.WorkloadSelector
"rules": []istio.security.v1beta1.Rule

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `selector` | [.istio.type.v1beta1.WorkloadSelector](../../../type/v1beta1/selector.proto.sk/#workloadselector) | Optional. Workload selector decides where to apply the authorization policy. If not set, the authorization policy will be applied to all workloads in the same namespace as the authorization policy. |  |
| `rules` | [[]istio.security.v1beta1.Rule](../authorization.proto.sk/#rule) | Optional. A list of rules to specify the allowed access to the workload. If not set, access is denied unless explicitly allowed by other authorization policy. |  |




---
### Rule

 
Rule allows access from a list of sources to perform a list of operations when
the condition is matched.

Any string field in the rule supports Exact, Prefix, Suffix and Presence match:
- Exact match: "abc" will match on value "abc".
- Prefix match: "abc*" will match on value "abc" and "abcd".
- Suffix match: "*abc" will match on value "abc" and "xabc".
- Presence match: "*" will match when value is not empty.

```yaml
"from": []istio.security.v1beta1.Rule.From
"to": []istio.security.v1beta1.Rule.To
"when": []istio.security.v1beta1.Condition

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `from` | [[]istio.security.v1beta1.Rule.From](../authorization.proto.sk/#from) | Optional. from specifies the source of a request. If not set, any source is allowed. |  |
| `to` | [[]istio.security.v1beta1.Rule.To](../authorization.proto.sk/#to) | Optional. to specifies the operation of a request. If not set, any operation is allowed. |  |
| `when` | [[]istio.security.v1beta1.Condition](../authorization.proto.sk/#condition) | Optional. when specifies a list of additional conditions of a request. If not set, any condition is allowed. |  |




---
### From

 
From includes a list or sources.

```yaml
"source": .istio.security.v1beta1.Source

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `source` | [.istio.security.v1beta1.Source](../authorization.proto.sk/#source) | Source specifies the source of a request. |  |




---
### To

 
To includes a list or operations.

```yaml
"operation": .istio.security.v1beta1.Operation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `operation` | [.istio.security.v1beta1.Operation](../authorization.proto.sk/#operation) | Operation specifies the operation of a request. |  |




---
### Source

 
Source specifies the source identities of a request.

```yaml
"principals": []string
"requestPrincipals": []string
"namespaces": []string
"ipBlocks": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `principals` | `[]string` | Optional. A list of source peer identities (i.e. service account), which matches to the "source.principal" attribute. If not set, any principal is allowed. |  |
| `requestPrincipals` | `[]string` | Optional. A list of request identities (i.e. "iss/sub" claims), which matches to the "request.auth.principal" attribute. If not set, any request principal is allowed. |  |
| `namespaces` | `[]string` | Optional. A list of namespaces, which matches to the "source.namespace" attribute. If not set, any namespace is allowed. |  |
| `ipBlocks` | `[]string` | Optional. A list of IP blocks, which matches to the "source.ip" attribute. Single IP (e.g. "1.2.3.4") and CIDR (e.g. "1.2.3.0/24") are supported. If not set, any IP is allowed. |  |




---
### Operation

 
Operation specifies the operations of a request.

```yaml
"hosts": []string
"ports": []string
"methods": []string
"paths": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `hosts` | `[]string` | Optional. A list of hosts, which matches to the "request.host" attribute. If not set, any host is allowed. Must be used only with HTTP. |  |
| `ports` | `[]string` | Optional. A list of ports, which matches to the "destination.port" attribute. If not set, any port is allowed. |  |
| `methods` | `[]string` | Optional. A list of methods, which matches to the "request.method" attribute. For gRPC service, this should be the fully-qualified name in the form of "/package.service/method" If not set, any method is allowed. Must be used only with HTTP or gRPC. |  |
| `paths` | `[]string` | Optional. A list of paths, which matches to the "request.url_path" attribute. If not set, any path is allowed. Must be used only with HTTP. |  |




---
### Condition

 
Condition specifies additional required attributes.

```yaml
"key": string
"values": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `key` | `string` | The name of an Istio attribute. See the [full list of supported attributes](https://istio.io/docs/reference/config/). |  |
| `values` | `[]string` | The allowed values for the attribute. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->