package endpoint

import "net/http"

// QueryParams parses whitelisted query parameters into the context.
func QueryParams(required []string, optional []string) Middleware {

	fn := func(ctx Context, h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			params := r.URL.Query()

			// required keys
			for _, key := range required {
				val := params.Get(key)
				if val == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				ctx[key] = val
			}

			// optional keys
			for _, key := range optional {
				val := params.Get(key)
				if val == "" {
					continue
				}
				ctx[key] = val
			}

			h.ServeHTTP(w, r)
		})
	}

	return fn
}
