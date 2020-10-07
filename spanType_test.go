package tracing_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/we-money/tracing"
	"testing"
)

func TestGetSpanType(t *testing.T) {
	urls := map[string]string{
		"health":   "/",
		"v1":       "/v1",
		"track":    "/v1/users/6a6b98e3-58b8-485d-b6f0-8556cc9f234a/track",
		"users":    "/v1/users/6a6b98e3-58b8-485d-b6f0-8556cc9f234a",
		"accounts": "/v1/users/6a6b98e3-58b8-485d-b6f0-8556cc9f234a/accounts/10203989",
		"balances": "v1/users/e2397294-28ea-4b56-a6a8-b18a1112b6d6/accounts/10203989/balances/2020-03-05/2020-09-10",
	}

	for k, v := range urls {
		assert.Equal(t, k, tracing.GetSpanType(v))
	}
}
