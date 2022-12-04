package durov

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComposeHandlers(t *testing.T) {
	middlewares := []func(Handler) Handler{
		func(handler Handler) Handler {
			return func(req *RequestContext) {
				req.CallbackData += "1"
				handler(req)
			}
		},
		func(handler Handler) Handler {
			return func(req *RequestContext) {
				req.CallbackData += "2"
				handler(req)
			}
		},
	}

	lastHandler := func(req *RequestContext) {
		req.CallbackData += "3"
	}

	handler := composeHandlers(middlewares, lastHandler)
	req := &RequestContext{}
	handler(req)

	assert.Equal(t, "123", req.CallbackData)
}
