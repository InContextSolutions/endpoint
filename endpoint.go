package endpoint

import (
	"fmt"
	"net/http"
)

// Context is maps strings to arbitrary types
type Context map[string]interface{}

// Controller is a function that accepts a context and returns a handler.
type Controller func(Context) http.Handler

// Middleware is a function that accepts a context and a handler and returns a handler.
type Middleware func(Context, http.Handler) http.Handler

// Endpoint is an endpoint on the server.
type Endpoint struct {
	Path    string       // Path is the is the url path.
	Method  string       // Method is the request method.
	Before  []Middleware // Middleware stack
	Control Controller   // Control handles the final portion of the request
}

// Handler joins the middleware with the controller.
func (e Endpoint) Handler() http.Handler {
	var mw []Middleware

	switch e.Method {
	case GET:
		mw = append(mw, get)
	case POST:
		mw = append(mw, post)
	default:
		err := fmt.Errorf("`%v` is not a supported http method", e.Method)
		panic(err)
	}

	mw = append(mw, e.Before...)
	return handler{before: mw, control: e.Control}
}
