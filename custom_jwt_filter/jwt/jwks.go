package jwt

import (
	"fmt"
	"strings"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

var host = "outbound|443||www.googleapis.com"

// var host = "googleapis"

// Callback represents callback function that has response result as arguments
type callback func(headers map[string]string, body []byte, err error)

func call(iss string, callback callback) error {
	cb := func(numHeaders, bodySize, numTrailers int) {
		hs, err := proxywasm.GetHttpCallResponseHeaders()
		if err != nil {
			callback(nil, nil, fmt.Errorf("failed to get response headers: %s", err))
			return
		}
		proxywasm.LogInfof("jwks response header %v", hs)
		respHeader := make(map[string]string)
		for _, v := range hs {
			respHeader[strings.TrimLeft(v[0], ":")] = v[1]
		}

		var body []byte
		if bodySize > 0 {
			body, err = proxywasm.GetHttpCallResponseBody(0, bodySize)
			proxywasm.LogInfof("jwks response body %v", string(body))
			if err != nil {
				callback(nil, nil, fmt.Errorf("failed to get response body: %s", err))
				return
			}
		}
		callback(respHeader, body, nil)
	}
	reqHeaders := [][2]string{
		{":authority", "www.googleapis.com"},
		{":method", "GET"},
		{":path", "/service_accounts/v1/metadata/jwk/" + iss},
	}
	if _, err := proxywasm.DispatchHttpCall(host, reqHeaders, "", nil, 5000, cb); err != nil {
		return err
	}
	return nil
}
