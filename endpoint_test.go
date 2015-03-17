package endpoint

import (
	"net/http"
	"testing"
)

func TestSmokeEndpoint(t *testing.T) {

	mw := func(ctx Context, h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			h(w, r)
		}
	}

	e := Endpoint{
		Path:         "/foo",
		Method:       "PUT",
		Before:       []Middleware{mw},
		RequiredArgs: []string{"param1"},
		OptionalArgs: []string{"param2"},
		Control: func(ctx Context) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
			}
		},
	}
	e.Handler()
}
