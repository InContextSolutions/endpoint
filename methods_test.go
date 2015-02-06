package endpoint

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ok struct {
}

func (o ok) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func TestGetWorks(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := get(ctx, ok{})
	g.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 200, "did not get status 200")
}

func TestGetFails(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := get(ctx, ok{})
	g.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 400, "did not get status 400")
}

func TestPostWorks(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, ok{})
	g.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 200, "did not get status 200")
}

func TestPostFails(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	ctx := make(Context)
	g := post(ctx, ok{})
	g.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 400, "did not get status 400")
}
