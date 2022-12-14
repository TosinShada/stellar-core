package test

import (
	"context"
	"testing"

	"github.com/TosinShada/stellar-core/exp/services/soroban-rpc/internal/methods"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	test := NewTest(t)

	ch := jhttp.NewChannel(test.server.URL, nil)
	cli := jrpc2.NewClient(ch, nil)

	var result methods.HealthCheckResult
	if err := cli.CallResult(context.Background(), "getHealth", nil, &result); err != nil {
		t.Fatalf("rpc call failed: %v", err)
	}
	assert.Equal(t, methods.HealthCheckResult{Status: "healthy"}, result)
}
