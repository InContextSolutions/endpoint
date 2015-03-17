package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryParams(t *testing.T) {
	q := queryParams([]string{"req"}, []string{"opt"})

	ctx := make(Context)
	h := func(w http.ResponseWriter, r *http.Request) {}
	handler := q(ctx, h)

	r1, _ := http.NewRequest("GET", "http://example.com/foo?req=1&opt=2", nil)
	handler(httptest.NewRecorder(), r1)

	r2, _ := http.NewRequest("GET", "http://example.com/foo?opt=2", nil)
	handler(httptest.NewRecorder(), r2)

	r3, _ := http.NewRequest("GET", "http://example.com/foo?req=1", nil)
	handler(httptest.NewRecorder(), r3)
}
