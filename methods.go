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
			body := json.NewDecoder(r.Body)
			err := body.Decode(&data)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ctx["data"] = data
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
