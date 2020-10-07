package tracing

import (
	"strings"
)

func GetSpanType(url string) string {
	if url == "/" {
		return "health"
	}

	u := strings.Split(url, "/")
	spanType := "yodler:handle"

	if len(u) == 2 {
		spanType = u[1]
	}
	if len(u) > 2 && len(u) <= 4 {
		spanType = u[2]
	}
	if len(u) > 4 && len(u) <= 6 {
		spanType = u[4]
	}
	if len(u) > 6 {
		spanType = u[5]
	}
	return spanType
}
