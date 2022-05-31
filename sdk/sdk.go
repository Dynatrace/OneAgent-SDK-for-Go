package sdk

import (
	"github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/internal"
	"github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/trace"
)

// OneAgentSDK provides the methods available in the OneAgent SDK for Go.
type OneAgentSDK interface {
	// GetTraceContextInfo returns information about the active PurePath node using
	// the TraceContext model as defined in https://www.w3.org/TR/trace-context.
	// The returned information is not intended for tagging and context-propagation
	// scenarios and primarily designed for log-enrichment use cases.
	GetTraceContextInfo() trace.TraceContextInfo

	// GetEnrichmentMetadata returns metadata that can be used to manually enrich
	// log messages when unsupported logging frameworks are used.
	// See also https://www.dynatrace.com/support/help/shortlink/enrich-metrics
	GetEnrichmentMetadata() map[string]string
}

// CreateInstance returns an instance of the OneAgent SDK.
//go:noinline
func CreateInstance() OneAgentSDK {
	return internal.OneAgentStubSDK{}
}
