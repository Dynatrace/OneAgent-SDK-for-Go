package trace

const (
	InvalidTraceId = "00000000000000000000000000000000"
	InvalidSpanId  = "0000000000000000"
)

// TraceContextInfo is the type returned by `OneAgentSDK.GetTraceContextInfo()`. It provides information about the
// active PurePath node using the TraceContext (Trace ID, Span ID) model as defined in https://www.w3.org/TR/trace-context.
//
// The Span ID represents the active PurePath node and is not intended for tagging or context propagation scenarios,
// it is primarily designed for log-enrichment use cases.
type TraceContextInfo struct {
	traceId string
	spanId  string
}

// GetTraceId returns Trace ID represented as lower-case hex-encoded string
// (see: https://tools.ietf.org/html/rfc4648#section-8).
func (c *TraceContextInfo) GetTraceId() string {
	return c.traceId
}

// GetSpanId returns Span ID represented as lower-case hex-encoded string
// (see: https://tools.ietf.org/html/rfc4648#section-8).
func (c *TraceContextInfo) GetSpanId() string {
	return c.spanId
}

// IsValid returns true if TraceContextInfo has valid Trace ID and Span ID.
func (c *TraceContextInfo) IsValid() bool {
	return c.traceId != InvalidTraceId && c.spanId != InvalidSpanId
}

// NewTraceContextInfo creates an instance of TraceContextInfo with given Trace ID and Span ID.
func NewTraceContextInfo(traceId, spanId string) TraceContextInfo {
	return TraceContextInfo{
		traceId: traceId,
		spanId:  spanId,
	}
}
