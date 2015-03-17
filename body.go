package endpoint

import (
	"io/ioutil"
	"net/http"
)

func readBody() Middleware {
	fn := func(ctx Context, h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("Content-Type") != "application/json" {
				w.Write([]byte(`{"message":"content type must be application/json"}`))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if r.ContentLength > 0 {
				if data, err := ioutil.ReadAll(r.Body); err == nil {
					ctx["data"] = data
					h(w, r)
					return
				}
			}

			w.Write([]byte(`{"message":"body is malformed or missing"}`))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	return fn
}
