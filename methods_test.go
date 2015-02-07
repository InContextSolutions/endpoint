package endpoint

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SayOK struct {
}

func (o SayOK) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func TestGetWorks(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := get(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
}

func TestGetBlocksPost(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := get(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 405, w.Code, "did not get status 405")
}

func TestPostWorks(t *testing.T) {
	d := []byte(`{"Answer": "42"}`)
	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewReader(d))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code, "did not get status 200")
}

func TestPostMissingContentType(t *testing.T) {
	d := []byte(`{"Answer": "42"}`)
	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewReader(d))
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code, "did not get status 400")
}

func TestPostWrongContentType(t *testing.T) {
	d := []byte(`{"Answer": "42"}`)
	r, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewReader(d))
	r.Header.Set("Content-Type", "gooba/gabba")
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code, "did not get status 400")
}

func TestEmptyPostFails(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 400, w.Code, "did not get status 400")
}

func TestPostBlocksGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, SayOK{})
	g.ServeHTTP(w, r)
	assert.Equal(t, 405, w.Code, "did not get status 405")
}
