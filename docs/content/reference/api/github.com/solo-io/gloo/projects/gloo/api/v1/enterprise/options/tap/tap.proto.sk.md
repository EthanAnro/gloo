
---
title: "tap.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `tap.options.gloo.solo.io` 
#### Types:


- [Tap](#tap)
- [Sink](#sink)
- [GrpcService](#grpcservice)
- [HttpService](#httpservice)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/tap/tap.proto](https://github.com/solo-io/gloo/blob/main/projects/gloo/api/v1/enterprise/options/tap/tap.proto)





---
### Tap

 
Tap filter: a filter that copies the contents of HTTP requests and responses
to an external tap server. The full HTTP headers and bodies are reported in
full to the configured address, and data can be reported using either over
HTTP or GRPC.

```yaml
"sinks": []tap.options.gloo.solo.io.Sink

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `sinks` | [[]tap.options.gloo.solo.io.Sink](../tap.proto.sk/#sink) | Sinks to which tap data should be output. Currently, only a single sink is supported. |




---
### Sink



```yaml
"grpcService": .tap.options.gloo.solo.io.GrpcService
"httpService": .tap.options.gloo.solo.io.HttpService

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `grpcService` | [.tap.options.gloo.solo.io.GrpcService](../tap.proto.sk/#grpcservice) | Write tap data out to a GRPC service. Only one of `grpcService` or `httpService` can be set. |
| `httpService` | [.tap.options.gloo.solo.io.HttpService](../tap.proto.sk/#httpservice) | Write tap data out to a HTTP service. Only one of `httpService` or `grpcService` can be set. |




---
### GrpcService

 
A tap sink over a GRPC service

```yaml
"tapServer": .core.solo.io.ResourceRef

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `tapServer` | [.core.solo.io.ResourceRef](../../../../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) | Upstream reference to the tap server. |




---
### HttpService

 
A tap sink over a HTTP service

```yaml
"tapServer": .core.solo.io.ResourceRef
"timeout": .google.protobuf.Duration

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `tapServer` | [.core.solo.io.ResourceRef](../../../../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) | Upstream reference to the tap server. |
| `timeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Connection timeout. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
