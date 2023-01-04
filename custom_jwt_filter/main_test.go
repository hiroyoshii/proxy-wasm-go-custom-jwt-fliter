package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func TestHttpAuthRandom_OnHttpRequestHeaders(t *testing.T) {
	opt := proxytest.NewEmulatorOption().
		WithNewRootContext(newRootContext)
	host := proxytest.NewHostEmulator(opt)
	// Release the host emulation lock so that other test cases can insert their own host emulation.
	defer host.Done()

	// Initialize context.
	contextID := host.InitializeHttpContext()

	// Call OnHttpRequestHeaders.
	action := host.CallOnRequestHeaders(contextID,
		types.Headers{{"key", "value"}}, false)
	require.Equal(t, types.ActionPause, action)

	// Verify DispatchHttpCall is called.
	attrs := host.GetCalloutAttributesFromContext(contextID)
	require.Equal(t, len(attrs), 1)
	require.Equal(t, "httpbin", attrs[0].Upstream)
	// Check if the current action is pause.
	require.Equal(t, types.ActionPause,
		host.GetCurrentHttpStreamAction(contextID))

	// Check Envoy logs.
	// logs := host.GetLogs(types.LogLevelInfo)
	// assert.Contains(t, logs, "http call dispatched to "+clusterName)
	// assert.Contains(t, logs, "request header: key: value")
}
