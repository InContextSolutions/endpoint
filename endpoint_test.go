package endpoint

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	e := Endpoint{
		Path:   "/foo",
		Method: GET,
		Before: []Middleware{addContext},
		Control: func(ctx Context) http.Handler {
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

	e.Handler().ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
	assert.Equal(t, "42", w.Body.String(), "did not get answer to the ultimate question")
}

func TestPostEndpoint(t *testing.T) {
	d := []byte(`{"Answer": "42"}`)

	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewReader(d))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	e := Endpoint{
		Path:   "/foo",
		Method: POST,
		Before: []Middleware{addContext},
		Control: func(ctx Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					data, haskey := ctx["data"]
					datab, _ := data.([]byte)
					if haskey {
						w.WriteHeader(http.StatusOK)
						w.Write(datab)
					} else {
						w.WriteHeader(http.StatusInternalServerError)
					}
				})
		},
	}

	e.Handler().ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
	assert.Equal(t, d, w.Body.Bytes(), "did not get answer to the ultimate question")
}

func TestPostEndpointBadJson(t *testing.T) {
	d := []byte(`{"Answer": "42",}`)

	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewReader(d))
	w := httptest.NewRecorder()

	e := Endpoint{
		Path:   "/foo",
		Method: POST,
		Before: []Middleware{addContext},
		Control: func(ctx Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
				})
		},
	}

	e.Handler().ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code, "did not get status 400")
}

func TestUnknownMethodEndpoint(t *testing.T) {
	e := Endpoint{
		Path:   "/foo",
		Method: "WHOOP",
		Before: []Middleware{addContext},
		Control: func(ctx Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				})
		},
	}

	assert.Panics(t, func() { e.Handler() }, "should panic")
}
