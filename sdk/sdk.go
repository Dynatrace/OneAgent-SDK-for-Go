// Package sdk provides the OneAgentSDK interface, which is the entry point to all OneAgent SDK for Go functionality.
// See also package sdk/native, which provides functionality that does not require the OneAgent Go code module.
//
// The first step in using the SDK is to acquire an instance of the OneAgentSDK interface by using the
// `CreateInstance` function.
package sdk

import (
	"github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/internal"
	"github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/trace"
)

// OneAgentSDK provides the methods available in the OneAgent SDK for Go. All methods are safe for concurrent use
// from multiple goroutines.
type OneAgentSDK interface {
	// GetTraceContextInfo returns information about the active PurePath node using the TraceContext model as defined
	// in https://www.w3.org/TR/trace-context. The returned information is not intended for tagging or
	// context propagation scenarios, it is primarily designed for log-enrichment use cases.
	//
	// The returned value may be invalid in case there is no active Dynatrace PurePath context.
	//
	// Since SDK version 1.0.0 - OneAgent version 1.233
	GetTraceContextInfo() trace.TraceContextInfo

	// GetEnrichmentMetadata returns metadata that can be used to manually enrich log messages when unsupported
	// logging frameworks are used.
	//
	// See also https://www.dynatrace.com/support/help/shortlink/enrich-metrics
	//
	// Since SDK version 1.1.0 - OneAgent version 1.245
	GetEnrichmentMetadata() map[string]string
}

// CreateInstance returns an instance of the OneAgent SDK.
//go:noinline
func CreateInstance() OneAgentSDK {
	return internal.OneAgentStubSDK{}
}
