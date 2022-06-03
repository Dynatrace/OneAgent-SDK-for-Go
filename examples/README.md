# Sample applications for OneAgent SDK for Go

Sample applications showing how to use Dynatrace OneAgent SDK for Go.

## Contents

* `trace_context`: Shows usage of the SDK to get Trace ID and Span ID information of the current PurePath node

## Prepare running sample applications

* Ensure Dynatrace OneAgent is installed. If not see our [free Trial](https://www.dynatrace.com/trial/)
* Ensure you have [the Go toolchain](https://golang.org "golang") installed
* Clone this repository
* Execute `go build` in the sample folder

## TraceContext sample application

This sample shows how to get Trace ID and Span ID information of the current PurePath node. Valid Trace ID and Span ID
can only be obtained if `GetTraceContextInfo` is executed in context of an active Dynatrace PurePath. As an example a
simple HTTP server is used, the agent will create a PurePath node which is active while running the HTTP handler. Make
sure that Go web request monitoring is enabled upon instrumentation of the sample application.

Execute the sample and send an HTTP GET request to `http://localhost:8080`, if everything is configured properly you
will get a response with valid Trace and Span IDs of the incoming HTTP request:

```text
Dynatrace TraceContext of the Incoming HTTP Request:
TraceId: 2c4a5bcd583ef3fa53ce939dbe0ea136, SpanId: 00f088ac0ba9c2b1, IsValid: true
```

Trace ID and Span ID values may be invalid in the following cases:
* OneAgent is not present
* OneAgent does not support the SDK version (check the [required version](#requirements))
* There is no active Dynatrace PurePath context
