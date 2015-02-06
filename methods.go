package endpoint

import (
	"encoding/json"
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
			var data map[string]interface{}
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				if err = json.Unmarshal(body, &data); err == nil {
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
