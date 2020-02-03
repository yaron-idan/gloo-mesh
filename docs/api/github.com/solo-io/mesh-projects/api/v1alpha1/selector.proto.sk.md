
---
title: "selector.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `zephyr.solo.io` 
#### Types:


- [PodSelector](#podselector)
- [LabelSelector](#labelselector)
- [ServiceSelector](#serviceselector)
- [NamespaceSelector](#namespaceselector)
  



##### Source File: [github.com/solo-io/mesh-projects/api/v1alpha1/selector.proto](https://github.com/solo-io/mesh-projects/blob/master/api/v1alpha1/selector.proto)





---
### PodSelector

 
specifies a method by which to select pods
with in a mesh for the application of rules and policies

```yaml
"labelSelector": .zephyr.solo.io.PodSelector.LabelSelector
"serviceSelector": .zephyr.solo.io.PodSelector.ServiceSelector
"namespaceSelector": .zephyr.solo.io.PodSelector.NamespaceSelector
"cluster": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `labelSelector` | [.zephyr.solo.io.PodSelector.LabelSelector](../../v1/selector.proto.sk/#labelselector) | select pods by their labels. Only one of `labelSelector`, or `namespaceSelector` can be set. |  |
| `serviceSelector` | [.zephyr.solo.io.PodSelector.ServiceSelector](../../v1/selector.proto.sk/#serviceselector) | select pods by their corresponding services. Only one of `serviceSelector`, or `namespaceSelector` can be set. |  |
| `namespaceSelector` | [.zephyr.solo.io.PodSelector.NamespaceSelector](../../v1/selector.proto.sk/#namespaceselector) | select all pods within one or more namespaces. Only one of `namespaceSelector`, or `serviceSelector` can be set. |  |
| `cluster` | `string` | cluster to search for the specified resources if not specified will default to the local cluster. |  |




---
### LabelSelector

 
select pods by their labels

```yaml
"labelsToMatch": map<string, string>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `labelsToMatch` | `map<string, string>` |  |  |




---
### ServiceSelector

 
select pods based on the services they back

```yaml
"services": []core.solo.io.ResourceRef

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `services` | [[]core.solo.io.ResourceRef](../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) | apply the selector to one or more services by adding their refs here. |  |




---
### NamespaceSelector

 
select all pods in these namespaces

```yaml
"namespaces": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `namespaces` | `[]string` |  |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->