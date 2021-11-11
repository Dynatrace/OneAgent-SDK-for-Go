package trace

const (
	InvalidTraceId = "00000000000000000000000000000000"
	InvalidSpanId  = "0000000000000000"
)

// TraceContextInfo provides information about the active PurePath node using
// the TraceContext (Trace ID, Span ID) model as defined in https://www.w3.org/TR/trace-context.
// The Span ID represents the active PurePath node.
// This Trace ID and Span ID information is not intended for tagging and context-propagation
// scenarios and primarily designed for log-enrichment use cases.
type TraceContextInfo struct {
	traceId string
	spanId  string
}

// GetTraceId returns Trace ID represented as lower-case hex-encoded string
// (see: https://tools.ietf.org/html/rfc4648#section-8).
// Returns an invalid Trace ID in case of:
// * OneAgent is not present.
// * OneAgent does not support the SDK version.
// * There is no active Dynatrace PurePath context.
func (c *TraceContextInfo) GetTraceId() string {
	return c.traceId
}

// GetSpanId returns Span ID represented as lower-case hex-encoded string
// (see: https://tools.ietf.org/html/rfc4648#section-8).
// Returns an invalid Span ID in case of:
// * OneAgent is not present.
// * OneAgent does not support the SDK version.
// * There is no active Dynatrace PurePath context.
func (c *TraceContextInfo) GetSpanId() string {
	return c.spanId
}

// IsValid returns true if TraceContextInfo has valid Trace ID and Span ID.
func (c *TraceContextInfo) IsValid() bool {
	return c.traceId != InvalidTraceId && c.spanId != InvalidSpanId
}

// NewTraceContextInfo create an instance of TraceContextInfo with given Trace ID and Span ID.
func NewTraceContextInfo(traceId, spanId string) TraceContextInfo {
	return TraceContextInfo{
		traceId: traceId,
		spanId:  spanId,
	}
}
