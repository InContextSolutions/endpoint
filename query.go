package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func queryParams(required []string, optional []string) Middleware {

	fn := func(ctx Context, h httprouter.Handle) httprouter.Handle {

		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

			qryParams := r.URL.Query()

			// required keys
			for _, key := range required {
				val := qryParams.Get(key)
				if val == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				ctx[key] = val
			}

			// optional keys
			for _, key := range optional {
				val := qryParams.Get(key)
				if val == "" {
					continue
				}
				ctx[key] = val
			}

			h(w, r, p)
		}
	}

	return fn
}
