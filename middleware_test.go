package endpoint

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	qp = QueryParams(
		[]string{"req1", "req2"},
		[]string{"opt1", "opt2"},
	)
	qe = Endpoint{
		Path:   "/foo",
		Method: GET,
		Before: []Middleware{qp},
		Control: func(ctx Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					req1 := ctx["req1"]
					req2 := ctx["req2"]
					opt1 := ctx["opt1"]
					opt2 := ctx["opt2"]
					dont := ctx["behere"]
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fmt.Sprintf("%v %v %v %v %v", req1, req2, opt1, opt2, dont)))
				})
		},
	}
)

func TestQueryParams(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo?req1=one&req2=two&opt1=three&opt2=four&dont=behere", nil)
	w := httptest.NewRecorder()
	qe.Handler().ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
	assert.Equal(t, "one two three four <nil>",
		w.Body.String(), "did not get correct query parameters")
}

func TestQueryParamsFail(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo?req2=two&opt1=three&opt2=four&dont=behere", nil)
	w := httptest.NewRecorder()
	qe.Handler().ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code, "did not get status 400")
}

func TestQueryParamsOptional(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo?req1=one&req2=two&opt2=four&dont=behere", nil)
	w := httptest.NewRecorder()
	qe.Handler().ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
	assert.Equal(t, "one two <nil> four <nil>",
		w.Body.String(), "did not get correct query parameters")
}
