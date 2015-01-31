package endpoint

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func addContext(ctx Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx["the answer"] = 42
		h.ServeHTTP(w, r)
	})
}

func TestNoMiddlewareInHandlerWorks(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	h := handler{
		before: []Middleware{},
		control: func(Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotImplemented)
				})
		},
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 501, "did not get status 501")
}

func TestOneMiddlewareInHandlerWorks(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	h := handler{
		before: []Middleware{get},
		control: func(Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotImplemented)
				})
		},
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 501, "did not get status 501")
}

func TestOneMiddlewareInHandlerFails(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	h := handler{
		before: []Middleware{post},
		control: func(Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotImplemented)
				})
		},
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 405, "did not get status 405")
}

func TestContextMiddleware(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	h := handler{
		before: []Middleware{get, addContext},
		control: func(ctx Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					answer, ok := ctx["the answer"]
					if ok {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte(fmt.Sprintf("%v", answer)))
					} else {
						w.WriteHeader(http.StatusInternalServerError)
					}
				})
		},
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 200, "did not get status 200")
	assert.Equal(t, w.Body.String(), "42", "did not get answer to the ultimate question")
}
