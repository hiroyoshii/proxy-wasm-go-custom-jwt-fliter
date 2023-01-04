package jwt

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

var alg = "RS256"

type jwt struct {
	alg string
	kid string
	exp int64
	iss string
}

func NewJwt(token string) *jwt {
	header, payload, err := extractJsonsFromJwt(token)
	if err != nil {
		panic(err)
	}
	alg, err := jsonparser.GetString(header, "alg")
	if err != nil {
		panic(err)
	}
	kid, err := jsonparser.GetString(header, "kid")
	if err != nil {
		panic(err)
	}
	exp, err := jsonparser.GetInt(payload, "exp")
	if err != nil {
		panic(err)
	}

	iss, err := jsonparser.GetString(payload, "iss")
	if err != nil {
		panic(err)
	}

	return &jwt{
		alg: alg,
		kid: kid,
		exp: exp,
		iss: iss,
	}
}

// verrify request JWT token
func (j *jwt) Verify() bool {
	proxywasm.LogDebugf("jwt parms: %v", j)
	return j.verifyExpiration() && j.verifyAlg() && j.verifyJwks()
}

func (j *jwt) verifyExpiration() bool {
	return time.Now().Unix() < j.exp
}
func (j *jwt) verifyAlg() bool {
	return j.alg == alg
}
func (j *jwt) verifyJwks() (valid bool) {
	err := call(j.iss, func(headers map[string]string, body []byte, err error) {
		if err != nil {
			panic(err)
		}
		// verify only kid existance
		jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			kid, err := jsonparser.GetString(value, "kid")
			proxywasm.LogDebugf("jwks kid %v", kid)
			if kid == j.kid {
				valid = true
			}
		}, "keys")
	})
	if err != nil {
		panic(err)
	}

	return valid
}

func extractJsonsFromJwt(token string) ([]byte, []byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("not jwt")
	}
	header, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, err
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, err
	}
	return header, payload, nil
}
