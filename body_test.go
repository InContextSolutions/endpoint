package endpoint

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmokeReadBody(t *testing.T) {
	r := readBody()

	ctx := make(Context)
	h := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}
	handler := r(ctx, h)

	r1, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	handler(httptest.NewRecorder(), r1, []httprouter.Param{})

	r2, _ := http.NewRequest("GET", "http://example.com/foo", bytes.NewBuffer([]byte("1")))
	handler(httptest.NewRecorder(), r2, []httprouter.Param{})

	r3, _ := http.NewRequest("GET", "http://example.com/foo", bytes.NewBuffer([]byte("1")))
	r3.Header.Add("Content-Type", "application/json")
	handler(httptest.NewRecorder(), r3, []httprouter.Param{})

	r4, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	r4.Header.Add("Content-Type", "application/json")
	handler(httptest.NewRecorder(), r4, []httprouter.Param{})
}
