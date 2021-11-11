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
}

// CreateInstance returns an instance of the OneAgent SDK.
//go:noinline
func CreateInstance() OneAgentSDK {
	return internal.OneAgentStubSDK{}
}
