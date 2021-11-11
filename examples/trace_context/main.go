package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dynatrace/OneAgent-SDK-for-Go/sdk"
)

func main() {
	// Create OneAgent SDK API instance
	var oneagentsdk = sdk.CreateInstance()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		// Get TraceContextInfo within the incoming HTTP request
		// to obtain Trace ID and Span ID of the active Dynatrace PurePath context
		traceContext := oneagentsdk.GetTraceContextInfo()

		fmt.Fprintf(w, "Dynatrace TraceContext of the Incoming HTTP Request:\n"+
			"TraceId: %s, SpanId: %s, IsValid: %t",
			traceContext.GetTraceId(), traceContext.GetSpanId(), traceContext.IsValid())
	})

	fmt.Println("Starting HTTP server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
