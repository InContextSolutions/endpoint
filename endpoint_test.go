package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func TestSmokeEndpoint(t *testing.T) {

	mw := func(ctx Context, h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			h(w, r, p)
		}
	}

	e := Endpoint{
		Path:         "/foo",
		Method:       "PUT",
		Before:       []Middleware{mw},
		RequiredArgs: []string{"param1"},
		OptionalArgs: []string{"param2"},
		Control: func(ctx Context) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			}
		},
	}
	e.Handler()
}
