
---
title: "descriptor.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `google.protobuf`  
Protocol Buffers - Google's data interchange format
Copyright 2008 Google Inc.  All rights reserved.
https://developers.google.com/protocol-buffers/

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.  
Author: kenton@google.com (Kenton Varda)
 Based on original Protocol Buffers design by
 Sanjay Ghemawat, Jeff Dean, and others.

The messages in this file describe the definitions found in .proto files.
A valid .proto file can be translated directly to a FileDescriptorProto
without any other information (e.g. without reading its imports).


 
##### Types:


- [FileDescriptorSet](#FileDescriptorSet)
- [FileDescriptorProto](#FileDescriptorProto)
- [DescriptorProto](#DescriptorProto)
- [ExtensionRange](#ExtensionRange)
- [ReservedRange](#ReservedRange)
- [FieldDescriptorProto](#FieldDescriptorProto)
- [Type](#Type)
- [Label](#Label)
- [OneofDescriptorProto](#OneofDescriptorProto)
- [EnumDescriptorProto](#EnumDescriptorProto)
- [EnumValueDescriptorProto](#EnumValueDescriptorProto)
- [ServiceDescriptorProto](#ServiceDescriptorProto)
- [MethodDescriptorProto](#MethodDescriptorProto)
- [FileOptions](#FileOptions)
- [OptimizeMode](#OptimizeMode)
- [MessageOptions](#MessageOptions)
- [FieldOptions](#FieldOptions)
- [CType](#CType)
- [JSType](#JSType)
- [OneofOptions](#OneofOptions)
- [EnumOptions](#EnumOptions)
- [EnumValueOptions](#EnumValueOptions)
- [ServiceOptions](#ServiceOptions)
- [MethodOptions](#MethodOptions)
- [IdempotencyLevel](#IdempotencyLevel)
- [UninterpretedOption](#UninterpretedOption)
- [NamePart](#NamePart)
- [SourceCodeInfo](#SourceCodeInfo)
- [Location](#Location)
- [GeneratedCodeInfo](#GeneratedCodeInfo)
- [Annotation](#Annotation)
  



##### Source File: `google/protobuf/descriptor.proto`





---
### <a name="FileDescriptorSet">FileDescriptorSet</a>

 
The protocol compiler can output a FileDescriptorSet containing the .proto
files it parses.

```yaml
"file": []google.protobuf.FileDescriptorProto

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `file` | [[]google.protobuf.FileDescriptorProto](../descriptor.proto.sk#FileDescriptorProto) |  |  |




---
### <a name="FileDescriptorProto">FileDescriptorProto</a>

 
Describes a complete .proto file.

```yaml
"name": string
"package": string
"dependency": []string
"publicDependency": []int
"weakDependency": []int
"messageType": []google.protobuf.DescriptorProto
"enumType": []google.protobuf.EnumDescriptorProto
"service": []google.protobuf.ServiceDescriptorProto
"extension": []google.protobuf.FieldDescriptorProto
"options": .google.protobuf.FileOptions
"sourceCodeInfo": .google.protobuf.SourceCodeInfo
"syntax": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `package` | `string` |  |  |
| `dependency` | `[]string` | Names of files imported by this file. |  |
| `publicDependency` | `[]int` | Indexes of the public imported files in the dependency list above. |  |
| `weakDependency` | `[]int` | Indexes of the weak imported files in the dependency list. For Google-internal migration only. Do not use. |  |
| `messageType` | [[]google.protobuf.DescriptorProto](../descriptor.proto.sk#DescriptorProto) | All top-level definitions in this file. |  |
| `enumType` | [[]google.protobuf.EnumDescriptorProto](../descriptor.proto.sk#EnumDescriptorProto) |  |  |
| `service` | [[]google.protobuf.ServiceDescriptorProto](../descriptor.proto.sk#ServiceDescriptorProto) |  |  |
| `extension` | [[]google.protobuf.FieldDescriptorProto](../descriptor.proto.sk#FieldDescriptorProto) |  |  |
| `options` | [.google.protobuf.FileOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/file-options) |  |  |
| `sourceCodeInfo` | [.google.protobuf.SourceCodeInfo](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/source-code-info) | This field contains optional information about the original source code. You may safely remove this entire field without harming runtime functionality of the descriptors -- the information is needed only by development tools. |  |
| `syntax` | `string` | The syntax of the proto file. The supported values are "proto2" and "proto3". |  |




---
### <a name="DescriptorProto">DescriptorProto</a>

 
Describes a message type.

```yaml
"name": string
"field": []google.protobuf.FieldDescriptorProto
"extension": []google.protobuf.FieldDescriptorProto
"nestedType": []google.protobuf.DescriptorProto
"enumType": []google.protobuf.EnumDescriptorProto
"extensionRange": []google.protobuf.DescriptorProto.ExtensionRange
"oneofDecl": []google.protobuf.OneofDescriptorProto
"options": .google.protobuf.MessageOptions
"reservedRange": []google.protobuf.DescriptorProto.ReservedRange
"reservedName": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `field` | [[]google.protobuf.FieldDescriptorProto](../descriptor.proto.sk#FieldDescriptorProto) |  |  |
| `extension` | [[]google.protobuf.FieldDescriptorProto](../descriptor.proto.sk#FieldDescriptorProto) |  |  |
| `nestedType` | [[]google.protobuf.DescriptorProto](../descriptor.proto.sk#DescriptorProto) |  |  |
| `enumType` | [[]google.protobuf.EnumDescriptorProto](../descriptor.proto.sk#EnumDescriptorProto) |  |  |
| `extensionRange` | [[]google.protobuf.DescriptorProto.ExtensionRange](../descriptor.proto.sk#ExtensionRange) |  |  |
| `oneofDecl` | [[]google.protobuf.OneofDescriptorProto](../descriptor.proto.sk#OneofDescriptorProto) |  |  |
| `options` | [.google.protobuf.MessageOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/message-options) |  |  |
| `reservedRange` | [[]google.protobuf.DescriptorProto.ReservedRange](../descriptor.proto.sk#ReservedRange) |  |  |
| `reservedName` | `[]string` | Reserved field names, which may not be used by fields in the same message. A given name may only be reserved once. |  |




---
### <a name="ExtensionRange">ExtensionRange</a>



```yaml
"start": int
"end": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `start` | `int` |  |  |
| `end` | `int` |  |  |




---
### <a name="ReservedRange">ReservedRange</a>

 
Range of reserved tag numbers. Reserved tag numbers may not be used by
fields or extension ranges in the same message. Reserved ranges may
not overlap.

```yaml
"start": int
"end": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `start` | `int` |  |  |
| `end` | `int` |  |  |




---
### <a name="FieldDescriptorProto">FieldDescriptorProto</a>

 
Describes a field within a message.

```yaml
"name": string
"number": int
"label": .google.protobuf.FieldDescriptorProto.Label
"type": .google.protobuf.FieldDescriptorProto.Type
"typeName": string
"extendee": string
"defaultValue": string
"oneofIndex": int
"jsonName": string
"options": .google.protobuf.FieldOptions

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `number` | `int` |  |  |
| `label` | [.google.protobuf.FieldDescriptorProto.Label](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-descriptor-proto.-label) |  |  |
| `type` | [.google.protobuf.FieldDescriptorProto.Type](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-descriptor-proto.-type) | If type_name is set, this need not be set. If both this and type_name are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP. |  |
| `typeName` | `string` | For message and enum types, this is the name of the type. If the name starts with a '.', it is fully-qualified. Otherwise, C++-like scoping rules are used to find the type (i.e. first the nested types within this message are searched, then within the parent, on up to the root namespace). |  |
| `extendee` | `string` | For extensions, this is the name of the type being extended. It is resolved in the same manner as type_name. |  |
| `defaultValue` | `string` | For numeric types, contains the original text representation of the value. For booleans, "true" or "false". For strings, contains the default text contents (not escaped in any way). For bytes, contains the C escaped value. All bytes >= 128 are escaped. TODO(kenton): Base-64 encode? |  |
| `oneofIndex` | `int` | If set, gives the index of a oneof in the containing type's oneof_decl list. This field is a member of that oneof. |  |
| `jsonName` | `string` | JSON name of this field. The value is set by protocol compiler. If the user has set a "json_name" option on this field, that option's value will be used. Otherwise, it's deduced from the field's name by converting it to camelCase. |  |
| `options` | [.google.protobuf.FieldOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-options) |  |  |




---
### <a name="Type">Type</a>



| Name | Description |
| ----- | ----------- | 
| `TYPE_DOUBLE` | 0 is reserved for errors. Order is weird for historical reasons. |
| `TYPE_FLOAT` |  |
| `TYPE_INT64` | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT64 if negative values are likely. |
| `TYPE_UINT64` |  |
| `TYPE_INT32` | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT32 if negative values are likely. |
| `TYPE_FIXED64` |  |
| `TYPE_FIXED32` |  |
| `TYPE_BOOL` |  |
| `TYPE_STRING` |  |
| `TYPE_GROUP` | Tag-delimited aggregate. Group type is deprecated and not supported in proto3. However, Proto3 implementations should still be able to parse the group wire format and treat group fields as unknown fields. |
| `TYPE_MESSAGE` |  |
| `TYPE_BYTES` | New in version 2. |
| `TYPE_UINT32` |  |
| `TYPE_ENUM` |  |
| `TYPE_SFIXED32` |  |
| `TYPE_SFIXED64` |  |
| `TYPE_SINT32` |  |
| `TYPE_SINT64` |  |




---
### <a name="Label">Label</a>



| Name | Description |
| ----- | ----------- | 
| `LABEL_OPTIONAL` | 0 is reserved for errors |
| `LABEL_REQUIRED` |  |
| `LABEL_REPEATED` |  |




---
### <a name="OneofDescriptorProto">OneofDescriptorProto</a>

 
Describes a oneof.

```yaml
"name": string
"options": .google.protobuf.OneofOptions

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `options` | [.google.protobuf.OneofOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/oneof-options) |  |  |




---
### <a name="EnumDescriptorProto">EnumDescriptorProto</a>

 
Describes an enum type.

```yaml
"name": string
"value": []google.protobuf.EnumValueDescriptorProto
"options": .google.protobuf.EnumOptions

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `value` | [[]google.protobuf.EnumValueDescriptorProto](../descriptor.proto.sk#EnumValueDescriptorProto) |  |  |
| `options` | [.google.protobuf.EnumOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/enum-options) |  |  |




---
### <a name="EnumValueDescriptorProto">EnumValueDescriptorProto</a>

 
Describes a value within an enum.

```yaml
"name": string
"number": int
"options": .google.protobuf.EnumValueOptions

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `number` | `int` |  |  |
| `options` | [.google.protobuf.EnumValueOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/enum-value-options) |  |  |




---
### <a name="ServiceDescriptorProto">ServiceDescriptorProto</a>

 
Describes a service.

```yaml
"name": string
"method": []google.protobuf.MethodDescriptorProto
"options": .google.protobuf.ServiceOptions

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `method` | [[]google.protobuf.MethodDescriptorProto](../descriptor.proto.sk#MethodDescriptorProto) |  |  |
| `options` | [.google.protobuf.ServiceOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/service-options) |  |  |




---
### <a name="MethodDescriptorProto">MethodDescriptorProto</a>

 
Describes a method of a service.

```yaml
"name": string
"inputType": string
"outputType": string
"options": .google.protobuf.MethodOptions
"clientStreaming": bool
"serverStreaming": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` |  |  |
| `inputType` | `string` | Input and output type names. These are resolved in the same way as FieldDescriptorProto.type_name, but must refer to a message type. |  |
| `outputType` | `string` |  |  |
| `options` | [.google.protobuf.MethodOptions](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/method-options) |  |  |
| `clientStreaming` | `bool` | Identifies if client streams multiple client messages |  Default: false |
| `serverStreaming` | `bool` | Identifies if server streams multiple server messages |  Default: false |




---
### <a name="FileOptions">FileOptions</a>



```yaml
"javaPackage": string
"javaOuterClassname": string
"javaMultipleFiles": bool
"javaGenerateEqualsAndHash": bool
"javaStringCheckUtf8": bool
"optimizeFor": .google.protobuf.FileOptions.OptimizeMode
"goPackage": string
"ccGenericServices": bool
"javaGenericServices": bool
"pyGenericServices": bool
"deprecated": bool
"ccEnableArenas": bool
"objcClassPrefix": string
"csharpNamespace": string
"swiftPrefix": string
"phpClassPrefix": string
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `javaPackage` | `string` | Sets the Java package where classes generated from this .proto will be placed. By default, the proto package is used, but this is often inappropriate because proto packages do not normally start with backwards domain names. |  |
| `javaOuterClassname` | `string` | If set, all the classes from the .proto file are wrapped in a single outer class with the given name. This applies to both Proto1 (equivalent to the old "--one_java_file" option) and Proto2 (where a .proto always translates to a single class, but you may want to explicitly choose the class name). |  |
| `javaMultipleFiles` | `bool` | If set true, then the Java code generator will generate a separate .java file for each top-level message, enum, and service defined in the .proto file. Thus, these types will *not* be nested inside the outer class named by java_outer_classname. However, the outer class will still be generated to contain the file's getDescriptor() method as well as any top-level extensions defined in the file. |  Default: false |
| `javaGenerateEqualsAndHash` | `bool` | This option does nothing. |  |
| `javaStringCheckUtf8` | `bool` | If set true, then the Java2 code generator will generate code that throws an exception whenever an attempt is made to assign a non-UTF-8 byte sequence to a string field. Message reflection will do the same. However, an extension field still accepts non-UTF-8 byte sequences. This option has no effect on when used with the lite runtime. |  Default: false |
| `optimizeFor` | [.google.protobuf.FileOptions.OptimizeMode](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/file-options.-optimize-mode) |  |  Default: SPEED |
| `goPackage` | `string` | Sets the Go package where structs generated from this .proto will be placed. If omitted, the Go package will be derived from the following: - The basename of the package import path, if provided. - Otherwise, the package statement in the .proto file, if present. - Otherwise, the basename of the .proto file, without extension. |  |
| `ccGenericServices` | `bool` | Should generic services be generated in each language? "Generic" services are not specific to any particular RPC system. They are generated by the main code generators in each language (without additional plugins). Generic services were the only kind of service generation supported by early versions of google.protobuf. Generic services are now considered deprecated in favor of using plugins that generate code specific to your particular RPC system. Therefore, these default to false. Old code which depends on generic services should explicitly set them to true. |  Default: false |
| `javaGenericServices` | `bool` |  |  Default: false |
| `pyGenericServices` | `bool` |  |  Default: false |
| `deprecated` | `bool` | Is this file deprecated? Depending on the target platform, this can emit Deprecated annotations for everything in the file, or it will be completely ignored; in the very least, this is a formalization for deprecating files. |  Default: false |
| `ccEnableArenas` | `bool` | Enables the use of arenas for the proto messages in this file. This applies only to generated classes for C++. |  Default: false |
| `objcClassPrefix` | `string` | Sets the objective c class prefix which is prepended to all objective c generated classes from this .proto. There is no default. |  |
| `csharpNamespace` | `string` | Namespace for generated classes; defaults to the package. |  |
| `swiftPrefix` | `string` | By default Swift generators will take the proto package and CamelCase it replacing '.' with underscore and use that to prefix the types/symbols defined. When this options is provided, they will use this value instead to prefix the types/symbols defined. |  |
| `phpClassPrefix` | `string` | Sets the php class prefix which is prepended to all php generated classes from this .proto. Default is empty. |  |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="OptimizeMode">OptimizeMode</a>

 
Generated classes can be optimized for speed or code size.

| Name | Description |
| ----- | ----------- | 
| `SPEED` |  |
| `CODE_SIZE` | etc. |
| `LITE_RUNTIME` |  |




---
### <a name="MessageOptions">MessageOptions</a>



```yaml
"messageSetWireFormat": bool
"noStandardDescriptorAccessor": bool
"deprecated": bool
"mapEntry": bool
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `messageSetWireFormat` | `bool` | Set true to use the old proto1 MessageSet wire format for extensions. This is provided for backwards-compatibility with the MessageSet wire format. You should not use this for any other reason: It's less efficient, has fewer features, and is more complicated. The message must be defined exactly as follows: message Foo { option message_set_wire_format = true; extensions 4 to max; } Note that the message cannot have any defined fields; MessageSets only have extensions. All extensions of your type must be singular messages; e.g. they cannot be int32s, enums, or repeated messages. Because this is an option, the above two restrictions are not enforced by the protocol compiler. |  Default: false |
| `noStandardDescriptorAccessor` | `bool` | Disables the generation of the standard "descriptor()" accessor, which can conflict with a field of the same name. This is meant to make migration from proto1 easier; new code should avoid fields named "descriptor". |  Default: false |
| `deprecated` | `bool` | Is this message deprecated? Depending on the target platform, this can emit Deprecated annotations for the message, or it will be completely ignored; in the very least, this is a formalization for deprecating messages. |  Default: false |
| `mapEntry` | `bool` | Whether the message is an automatically generated map entry type for the maps field. For maps fields: map<KeyType, ValueType> map_field = 1; The parsed descriptor looks like: message MapFieldEntry { option map_entry = true; optional KeyType key = 1; optional ValueType value = 2; } repeated MapFieldEntry map_field = 1; Implementations may choose not to generate the map_entry=true message, but use a native map in the target language to hold the keys and values. The reflection APIs in such implementions still need to work as if the field is a repeated message field. NOTE: Do not set the option in .proto files. Always use the maps syntax instead. The option should only be implicitly set by the proto compiler parser. |  |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="FieldOptions">FieldOptions</a>



```yaml
"ctype": .google.protobuf.FieldOptions.CType
"packed": bool
"jstype": .google.protobuf.FieldOptions.JSType
"lazy": bool
"deprecated": bool
"weak": bool
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `ctype` | [.google.protobuf.FieldOptions.CType](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-options.c-type) | The ctype option instructs the C++ code generator to use a different representation of the field than it normally would. See the specific options below. This option is not yet implemented in the open source release -- sorry, we'll try to include it in a future version! |  Default: STRING |
| `packed` | `bool` | The packed option can be enabled for repeated primitive fields to enable a more efficient representation on the wire. Rather than repeatedly writing the tag and type for each element, the entire array is encoded as a single length-delimited blob. In proto3, only explicit setting it to false will avoid using packed encoding. |  |
| `jstype` | [.google.protobuf.FieldOptions.JSType](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-options.js-type) | The jstype option determines the JavaScript type used for values of the field. The option is permitted only for 64 bit integral and fixed types (int64, uint64, sint64, fixed64, sfixed64). By default these types are represented as JavaScript strings. This avoids loss of precision that can happen when a large value is converted to a floating point JavaScript numbers. Specifying JS_NUMBER for the jstype causes the generated JavaScript code to use the JavaScript "number" type instead of strings. This option is an enum to permit additional types to be added, e.g. goog.math.Integer. |  Default: JS_NORMAL |
| `lazy` | `bool` | Should this field be parsed lazily? Lazy applies only to message-type fields. It means that when the outer message is initially parsed, the inner message's contents will not be parsed but instead stored in encoded form. The inner message will actually be parsed when it is first accessed. This is only a hint. Implementations are free to choose whether to use eager or lazy parsing regardless of the value of this option. However, setting this option true suggests that the protocol author believes that using lazy parsing on this field is worth the additional bookkeeping overhead typically needed to implement it. This option does not affect the public interface of any generated code; all method signatures remain the same. Furthermore, thread-safety of the interface is not affected by this option; const methods remain safe to call from multiple threads concurrently, while non-const methods continue to require exclusive access. Note that implementations may choose not to check required fields within a lazy sub-message. That is, calling IsInitialized() on the outer message may return true even if the inner message has missing required fields. This is necessary because otherwise the inner message would have to be parsed in order to perform the check, defeating the purpose of lazy parsing. An implementation which chooses not to check required fields must be consistent about it. That is, for any particular sub-message, the implementation must either *always* check its required fields, or *never* check its required fields, regardless of whether or not the message has been parsed. |  Default: false |
| `deprecated` | `bool` | Is this field deprecated? Depending on the target platform, this can emit Deprecated annotations for accessors, or it will be completely ignored; in the very least, this is a formalization for deprecating fields. |  Default: false |
| `weak` | `bool` | For Google-internal migration only. Do not use. |  Default: false |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="CType">CType</a>



| Name | Description |
| ----- | ----------- | 
| `STRING` | Default mode. |
| `CORD` |  |
| `STRING_PIECE` |  |




---
### <a name="JSType">JSType</a>



| Name | Description |
| ----- | ----------- | 
| `JS_NORMAL` | Use the default type. |
| `JS_STRING` | Use JavaScript strings. |
| `JS_NUMBER` | Use JavaScript numbers. |




---
### <a name="OneofOptions">OneofOptions</a>



```yaml
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="EnumOptions">EnumOptions</a>



```yaml
"allowAlias": bool
"deprecated": bool
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `allowAlias` | `bool` | Set this option to true to allow mapping different tag names to the same value. |  |
| `deprecated` | `bool` | Is this enum deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum, or it will be completely ignored; in the very least, this is a formalization for deprecating enums. |  Default: false |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="EnumValueOptions">EnumValueOptions</a>



```yaml
"deprecated": bool
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `deprecated` | `bool` | Is this enum value deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum value, or it will be completely ignored; in the very least, this is a formalization for deprecating enum values. |  Default: false |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="ServiceOptions">ServiceOptions</a>



```yaml
"deprecated": bool
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `deprecated` | `bool` | Is this service deprecated? Depending on the target platform, this can emit Deprecated annotations for the service, or it will be completely ignored; in the very least, this is a formalization for deprecating services. |  Default: false |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="MethodOptions">MethodOptions</a>



```yaml
"deprecated": bool
"idempotencyLevel": .google.protobuf.MethodOptions.IdempotencyLevel
"uninterpretedOption": []google.protobuf.UninterpretedOption

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `deprecated` | `bool` | Is this method deprecated? Depending on the target platform, this can emit Deprecated annotations for the method, or it will be completely ignored; in the very least, this is a formalization for deprecating methods. |  Default: false |
| `idempotencyLevel` | [.google.protobuf.MethodOptions.IdempotencyLevel](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/method-options.-idempotency-level) |  |  Default: IDEMPOTENCY_UNKNOWN |
| `uninterpretedOption` | [[]google.protobuf.UninterpretedOption](../descriptor.proto.sk#UninterpretedOption) | The parser stores options it doesn't recognize here. See above. |  |




---
### <a name="IdempotencyLevel">IdempotencyLevel</a>

 
Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
or neither? HTTP based RPC implementation may choose GET verb for safe
methods, and PUT verb for idempotent methods instead of the default POST.

| Name | Description |
| ----- | ----------- | 
| `IDEMPOTENCY_UNKNOWN` |  |
| `NO_SIDE_EFFECTS` |  |
| `IDEMPOTENT` |  |




---
### <a name="UninterpretedOption">UninterpretedOption</a>

 
A message representing a option the parser does not recognize. This only
appears in options protos created by the compiler::Parser class.
DescriptorPool resolves these when building Descriptor objects. Therefore,
options protos in descriptor objects (e.g. returned by Descriptor::options(),
or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
in them.

```yaml
"name": []google.protobuf.UninterpretedOption.NamePart
"identifierValue": string
"positiveIntValue": int
"negativeIntValue": int
"doubleValue": float
"stringValue": bytes
"aggregateValue": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | [[]google.protobuf.UninterpretedOption.NamePart](../descriptor.proto.sk#NamePart) |  |  |
| `identifierValue` | `string` | The value of the uninterpreted option, in whatever type the tokenizer identified it as during parsing. Exactly one of these should be set. |  |
| `positiveIntValue` | `int` |  |  |
| `negativeIntValue` | `int` |  |  |
| `doubleValue` | `float` |  |  |
| `stringValue` | `bytes` |  |  |
| `aggregateValue` | `string` |  |  |




---
### <a name="NamePart">NamePart</a>

 
The name of the uninterpreted option.  Each string represents a segment in
a dot-separated name.  is_extension is true iff a segment represents an
extension (denoted with parentheses in options specs in .proto files).
E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
"foo.(bar.baz).qux".

```yaml
"namePart": string
"isExtension": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `namePart` | `string` |  |  |
| `isExtension` | `bool` |  |  |




---
### <a name="SourceCodeInfo">SourceCodeInfo</a>

 
Encapsulates information about the original source file from which a
FileDescriptorProto was generated.

```yaml
"location": []google.protobuf.SourceCodeInfo.Location

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `location` | [[]google.protobuf.SourceCodeInfo.Location](../descriptor.proto.sk#Location) | A Location identifies a piece of source code in a .proto file which corresponds to a particular definition. This information is intended to be useful to IDEs, code indexers, documentation generators, and similar tools. For example, say we have a file like: message Foo { optional string foo = 1; } Let's look at just the field definition: optional string foo = 1; ^ ^^ ^^ ^ ^^^ a bc de f ghi We have the following locations: span path represents [a,i) [ 4, 0, 2, 0 ] The whole field definition. [a,b) [ 4, 0, 2, 0, 4 ] The label (optional). [c,d) [ 4, 0, 2, 0, 5 ] The type (string). [e,f) [ 4, 0, 2, 0, 1 ] The name (foo). [g,h) [ 4, 0, 2, 0, 3 ] The number (1). Notes: - A location may refer to a repeated field itself (i.e. not to any particular index within it). This is used whenever a set of elements are logically enclosed in a single code segment. For example, an entire extend block (possibly containing multiple extension definitions) will have an outer location whose path refers to the "extensions" repeated field without an index. - Multiple locations may have the same path. This happens when a single logical declaration is spread out across multiple places. The most obvious example is the "extend" block again -- there may be multiple extend blocks in the same scope, each of which will have the same path. - A location's span is not always a subset of its parent's span. For example, the "extendee" of an extension declaration appears at the beginning of the "extend" block and is shared by all extensions within the block. - Just because a location's span is a subset of some other location's span does not mean that it is a descendent. For example, a "group" defines both a type and a field in a single declaration. Thus, the locations corresponding to the type and field and their components will overlap. - Code which tries to interpret locations should probably be designed to ignore those that it doesn't understand, as more types of locations could be recorded in the future. |  |




---
### <a name="Location">Location</a>



```yaml
"path": []int
"span": []int
"leadingComments": string
"trailingComments": string
"leadingDetachedComments": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `path` | `[]int` | Identifies which part of the FileDescriptorProto was defined at this location. Each element is a field number or an index. They form a path from the root FileDescriptorProto to the place where the definition. For example, this path: [ 4, 3, 2, 7, 1 ] refers to: file.message_type(3) // 4, 3 .field(7) // 2, 7 .name() // 1 This is because FileDescriptorProto.message_type has field number 4: repeated DescriptorProto message_type = 4; and DescriptorProto.field has field number 2: repeated FieldDescriptorProto field = 2; and FieldDescriptorProto.name has field number 1: optional string name = 1; Thus, the above path gives the location of a field name. If we removed the last element: [ 4, 3, 2, 7 ] this path refers to the whole field declaration (from the beginning of the label to the terminating semicolon). |  |
| `span` | `[]int` | Always has exactly three or four elements: start line, start column, end line (optional, otherwise assumed same as start line), end column. These are packed into a single field for efficiency. Note that line and column numbers are zero-based -- typically you will want to add 1 to each before displaying to a user. |  |
| `leadingComments` | `string` | If this SourceCodeInfo represents a complete declaration, these are any comments appearing before and after the declaration which appear to be attached to the declaration. A series of line comments appearing on consecutive lines, with no other tokens appearing on those lines, will be treated as a single comment. leading_detached_comments will keep paragraphs of comments that appear before (but not connected to) the current element. Each paragraph, separated by empty lines, will be one comment element in the repeated field. Only the comment content is provided; comment markers (e.g. //) are stripped out. For block comments, leading whitespace and an asterisk will be stripped from the beginning of each line other than the first. Newlines are included in the output. Examples: optional int32 foo = 1; // Comment attached to foo. // Comment attached to bar. optional int32 bar = 2; optional string baz = 3; // Comment attached to baz. // Another line attached to baz. // Comment attached to qux. // // Another line attached to qux. optional double qux = 4; // Detached comment for corge. This is not leading or trailing comments // to qux or corge because there are blank lines separating it from // both. // Detached comment for corge paragraph 2. optional string corge = 5; /* Block comment attached * to corge. Leading asterisks * will be removed. */ /* Block comment attached to * grault. */ optional int32 grault = 6; // ignored detached comments. |  |
| `trailingComments` | `string` |  |  |
| `leadingDetachedComments` | `[]string` |  |  |




---
### <a name="GeneratedCodeInfo">GeneratedCodeInfo</a>

 
Describes the relationship between generated code and its original source
file. A GeneratedCodeInfo message is associated with only one generated
source file, but may contain references to different source .proto files.

```yaml
"annotation": []google.protobuf.GeneratedCodeInfo.Annotation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `annotation` | [[]google.protobuf.GeneratedCodeInfo.Annotation](../descriptor.proto.sk#Annotation) | An Annotation connects some span of text in generated code to an element of its generating .proto file. |  |




---
### <a name="Annotation">Annotation</a>



```yaml
"path": []int
"sourceFile": string
"begin": int
"end": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `path` | `[]int` | Identifies the element in the original source .proto file. This field is formatted the same as SourceCodeInfo.Location.path. |  |
| `sourceFile` | `string` | Identifies the filesystem path to the original source .proto. |  |
| `begin` | `int` | Identifies the starting offset in bytes in the generated code that relates to the identified object. |  |
| `end` | `int` | Identifies the ending offset in bytes in the generated code that relates to the identified offset. The end offset should be one past the last relevant byte (so the length of the text = end - begin). |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->