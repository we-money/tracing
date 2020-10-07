package tracing

import (
	"net/http"

	"github.com/google/uuid"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

const TraceID = "traceid"

func CustomTraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var traceId apm.TraceID
		if strings := r.Header[apmhttp.W3CTraceparentHeader]; len(strings) == 1 && strings[0] != "" {
			if traceContext, err := apmhttp.ParseTraceparentHeader(strings[0]); err == nil {
				traceId = traceContext.Trace
			}
		}
		if err := traceId.Validate(); err != nil {
			var spanId apm.SpanID
			var traceOptions apm.TraceOptions
			u := uuid.New()
			copy(traceId[:], u[:])
			copy(spanId[:], traceId[8:])
			traceContext := apm.TraceContext{
				Trace:   traceId,
				Span:    spanId,
				Options: traceOptions.WithRecorded(true),
			}
			r.Header.Set(apmhttp.W3CTraceparentHeader, apmhttp.FormatTraceparentHeader(traceContext))
		}

		w.Header().Set(TraceID, traceId.String())
		next.ServeHTTP(w, r)
	})
}
