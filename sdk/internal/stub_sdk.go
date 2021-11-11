package internal

import "github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/trace"

type OneAgentStubSDK struct{}

// GetTraceContextInfo returns an invalid TraceContextInfo instance.
// This method will be replaced by OneAgent to provide the actual values.
//go:noinline
func (OneAgentStubSDK) GetTraceContextInfo() trace.TraceContextInfo {
	return trace.NewTraceContextInfo(trace.InvalidTraceId, trace.InvalidSpanId)
}
