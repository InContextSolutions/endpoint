package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Context is maps strings to arbitrary types
type Context map[string]interface{}

// Controller is a function that accepts a context and returns a handle.
type Controller func(Context) httprouter.Handle

// Middleware is a function that accepts a context and a handler and returns a handle.
type Middleware func(Context, httprouter.Handle) httprouter.Handle

// Endpoint is an endpoint on the server.
type Endpoint struct {
	Path         string
	Method       string
	Before       []Middleware
	RequiredArgs []string
	OptionalArgs []string
	Control      Controller
}

// Handler joins the middleware with the controller.
func (e Endpoint) Handler() httprouter.Handle {

	fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		ctx := make(Context)
		final := e.Control(ctx)

		// parse out query params
		if len(e.RequiredArgs) > 0 || len(e.OptionalArgs) > 0 {
			final = queryParams(e.RequiredArgs, e.OptionalArgs)(ctx, final)
		}

		// read body for PUT & POST
		if e.Method == "PUT" || e.Method == "POST" || e.Method == "PATCH" {
			final = readBody()(ctx, final)
		}

		for i := len(e.Before) - 1; i >= 0; i-- {
			final = e.Before[i](ctx, final)
		}

		final(w, r, p)
	}
	return fn
}
