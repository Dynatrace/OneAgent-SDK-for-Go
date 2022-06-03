package internal

import "github.com/Dynatrace/OneAgent-SDK-for-Go/sdk/trace"

// The methods of this type will be replaced by OneAgent to provide their actual return values.
type OneAgentStubSDK struct{}

//go:noinline
func (OneAgentStubSDK) GetTraceContextInfo() trace.TraceContextInfo {
	return trace.NewTraceContextInfo(trace.InvalidTraceId, trace.InvalidSpanId)
}

//go:noinline
func (OneAgentStubSDK) GetEnrichmentMetadata() map[string]string {
	return nil
}
