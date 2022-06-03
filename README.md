# Dynatrace OneAgent SDK for Go

OneAgent SDK for Go has to be used together with OneAgent injected. Otherwise, the dummy implementation of OneAgent SDK
will be used. The current implementation of OneAgent SDK for Go is made with a single purpose: providing access to the
Trace ID and Span ID information of the PurePath node.
This information can then be used, for example, to provide additional context in log messages.

## Table of Contents

* [Package contents](#package-contents)
* [Requirements](#requirements)
* [Integration](#integration)
  * [Troubleshooting](#troubleshooting)
* [API Concepts](#api-concepts)
  * [OneAgentSDK object](#oneagentsdk-object)
  * [Trace Context](#trace-context)
* [Help & Support](#help--support)
* [Release notes](#release-notes)

## Package contents

* `examples`: contains sample application, which demonstrates the usage of the SDK. See readme inside the directory
  for more details.
* `LICENSE`: license under which the whole SDK and sample applications are published.

## Requirements

* Deep monitoring by Dynatrace OneAgent must be successfully activated.

| OneAgent SDK for Go | Minimum OneAgent version | Support status |
|:--------------------|:-------------------------|:---------------|
| 1.0.0               | 1.233                    | Supported      |
| 0.1.0               | 1.233                    | Not supported  |

## Integration

Using this module should not cause any errors if OneAgent is not present (e.g. in testing) since the stub implementation
of OneAgent SDK for Go will be used.

### Troubleshooting

If the SDK can not connect to OneAgent, verify that a matching version of OneAgent is used. Check if OneAgent Go log
file contains the following log message: `OneAgent SDK has been successfully resolved`.

## API Concepts

Common concepts of the Dynatrace OneAgent SDK are explained in the
[Dynatrace OneAgent SDK repository](https://github.com/Dynatrace/OneAgent-SDK).

### OneAgentSDK object

The first step is to acquire an instance of the OneAgent SDK API by calling `sdk.CreateInstance()`.

```go
import "github.com/Dynatrace/OneAgent-SDK-for-Go/sdk"
...
oneagentsdk := sdk.CreateInstance()
```

### Trace Context

The obtained OneAgent SDK API instance provides the `GetTraceContextInfo()` method, which returns a `TraceContextInfo`
object that provides the Trace ID and Span ID of the current PurePath node.
`TraceContextInfo.IsValid()` may be used to verify that the Trace ID and Span ID are both valid (non-zero).
Trace ID and Span ID information is not intended for tagging and context-propagation scenarios and primarily designed
for log-enrichment use cases.

```go
...
traceContext := oneagentsdk.GetTraceContextInfo()
traceId := traceContext.getTraceId()
spanId := traceContext.getSpanId()
```

Trace ID and Span ID values may be invalid in the following cases:
* OneAgent is not present
* OneAgent does not support the SDK version (check the [required version](#requirements))
* There is no active Dynatrace PurePath context

## Help & Support

**Support policy**

Dynatrace OneAgent SDK for Go has alpha status. The features are currently being tested by Dynatrace and subject to change.

For a detailed support policy see [Dynatrace OneAgent SDK help](https://github.com/Dynatrace/OneAgent-SDK#help).

### Get Help

* Ask a question in the [product forums](https://answers.dynatrace.com/spaces/482/view.html)
* Read the [product documentation](https://www.dynatrace.com/support/help/)

### Open a [GitHub issue](https://github.com/Dynatrace/OneAgent-SDK-for-Go/issues) to

* Report minor defects, minor items or typos
* Ask for improvements or changes in the SDK API
* Ask any questions related to the community effort

SLAs don't apply for GitHub tickets.

### Customers can open a ticket on the [Dynatrace support portal](https://support.dynatrace.com/supportportal/) to

* Get support from the Dynatrace technical support engineering team
* Manage and resolve product related technical issues

SLAs apply according to the customer's support level.

## Release Notes

see also [Releases](https://github.com/Dynatrace/OneAgent-SDK-for-Go/releases)

| Version | Description                                |
|:--------|:-------------------------------------------|
| 0.1.0   | Initial Alpha Release                      |
