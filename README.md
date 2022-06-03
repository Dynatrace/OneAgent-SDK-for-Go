# Dynatrace OneAgent SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/Dynatrace/OneAgent-SDK-for-Go.svg)](https://pkg.go.dev/github.com/Dynatrace/OneAgent-SDK-for-Go)

This SDK provides functionality for Go applications to enrich their log messages and OpenTelemetry traces with metadata
useful in a Dynatrace environment.


## Package contents

* `examples`: contains sample applications, which demonstrate usage of the SDK. See readme inside the directory
  for more details.
* `LICENSE`: license under which the whole SDK and sample applications are published.


## Requirements

* Some features require the OneAgent Go code module to be injected successfully (i.e. deep monitoring being activated),
  while other features work without it. OneAgent is required in any case.

| OneAgent SDK for Go | Minimum OneAgent version | Support status |
|:--------------------|:-------------------------|:---------------|
| 1.1.0               | 1.245                    | Supported      |
| 1.0.0               | 1.233                    | Supported      |
| 0.1.0               | 1.233                    | Not supported  |

### Integration

Using the functions in this module without OneAgent present should be perfectly fine. The functions will then return
invalid values or an error. This may happen in the following cases:

* OneAgent is not present.
* OneAgent does not support the SDK version (check the [required version](#requirements)).
* Some other reason specific to a particular function.

### Troubleshooting

If the SDK can not connect to OneAgent, verify that a matching version of OneAgent is used and deep monitoring is
activated. Check that the OneAgent Go log file contains the following log message:

```text
OneAgent SDK has been resolved successfully
```


## Features

Common API concepts of the Dynatrace OneAgent SDK are explained in the [Dynatrace OneAgent SDK repository][dt-sdk]. An
overview of features in this SDK will follow, see the [`sdk` package documentation][pkg-sdk] for detailed information.

| Feature                                | SDK Version |
|:---------------------------------------|:------------|
| [Enrichment metadata][func-enrichment] | ≥ 1.1.0     |
| [TraceContext][pkg-tracecontext]       | ≥ 1.0.0     |

[dt-sdk]: https://github.com/Dynatrace/OneAgent-SDK
[pkg-sdk]: https://pkg.go.dev/github.com/Dynatrace/OneAgent-SDK-for-Go/sdk
[pkg-tracecontext]: https://pkg.go.dev/github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/trace#TraceContextInfo
[iface-sdk]: https://pkg.go.dev/github.com/Dynatrace/OneAgent-SDK-for-Go/sdk#OneAgentSDK
[func-enrichment]: https://pkg.go.dev/github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/native#GetEnrichmentMetadata


## Help & Support

**Support policy**

Starting with version 1.0.0, Dynatrace OneAgent SDK for Go has GA status. The features are fully supported by Dynatrace.

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

See also [GitHub Releases](https://github.com/Dynatrace/OneAgent-SDK-for-Go/releases).

| Version | Description                                |
|:--------|:-------------------------------------------|
| 1.1.0   | Added support for enrichment metadata      |
| 1.0.0   | Initial GA release                         |
| 0.1.0   | Initial Alpha release                      |
