package main

import (
	"main/jwt"
	"strings"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

// Other examples can be found at https://github.com/tetratelabs/proxy-wasm-go-sdk/tree/v0.1.1/examples

func main() {
	proxywasm.SetNewRootContext(newRootContext)
}

type httpHeaders struct {
	// we must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultHttpContext
	contextID uint32
}

type rootContext struct {
	// You'd better embed the default root context
	// so that you don't need to reimplement all the methods by yourself.
	proxywasm.DefaultRootContext
}

func newRootContext(uint32) proxywasm.RootContext { return &rootContext{} }

// Override DefaultRootContext.
func (*rootContext) NewHttpContext(contextID uint32) proxywasm.HttpContext {
	return &httpHeaders{contextID: contextID}
}

// Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
	}
	proxywasm.LogDebugf("on request header: %v", hs)

	for _, h := range hs {
		// jwt header access
		if h[0] == "authorization" {
			token := strings.TrimPrefix(h[1], "Bearer ")
			if jwt.NewJwt(token).Verify() {
				return types.ActionContinue
			}
		}
	}
	return types.ActionPause
}

// Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	if err := proxywasm.SetHttpResponseHeader("hello", "world"); err != nil {
		proxywasm.LogCriticalf("failed to set response header: %v", err)
	}
	return types.ActionContinue
}

// Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpStreamDone() {
	proxywasm.LogInfof("%d finished", ctx.contextID)
}
