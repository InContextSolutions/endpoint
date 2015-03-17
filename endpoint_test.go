package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
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

	h := e.Handler()
	r, _ := http.NewRequest("PUT", "http://example.com/foo?param1=1&param2=2", nil)
	h(httptest.NewRecorder(), r, nil)
}

func TestCleanContext(t *testing.T) {

	mw := func(ctx Context, h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if _, hasKey := ctx["key"]; hasKey {
				t.Fail()
			} else {
				ctx["key"] = 1
			}
			h(w, r, p)
		}
	}

	e := Endpoint{
		Path:   "/foo",
		Method: "GET",
		Before: []Middleware{mw},
		Control: func(ctx Context) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			}
		},
	}

	h := e.Handler()

	r1, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	h(httptest.NewRecorder(), r1, nil)

	r2, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	h(httptest.NewRecorder(), r2, nil)
}
