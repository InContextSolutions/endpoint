package endpoint

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmokeReadBody(t *testing.T) {
	r := readBody()

	ctx := make(Context)
	h := func(w http.ResponseWriter, r *http.Request) {}
	handler := r(ctx, h)

	r1, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	handler(httptest.NewRecorder(), r1)

	r2, _ := http.NewRequest("GET", "http://example.com/foo", bytes.NewBuffer([]byte("1")))
	handler(httptest.NewRecorder(), r2)

	r3, _ := http.NewRequest("GET", "http://example.com/foo", bytes.NewBuffer([]byte("1")))
	r3.Header.Add("Content-Type", "application/json")
	handler(httptest.NewRecorder(), r3)

	r4, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	r4.Header.Add("Content-Type", "application/json")
	handler(httptest.NewRecorder(), r4)
}
