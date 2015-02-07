package endpoint

import (
	"io/ioutil"
	"net/http"
)

const (
	// GET request
	GET = "GET"

	// POST request
	POST = "POST"
)

func get(ctx Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == GET {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func post(ctx Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == POST {
			if r.ContentLength > 0 && r.Header.Get("Content-Type") == "application/json" {
				if data, err := ioutil.ReadAll(r.Body); err == nil {
					ctx["data"] = data
					h.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
