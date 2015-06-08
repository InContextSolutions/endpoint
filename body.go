package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func readBody() Middleware {
	fn := func(ctx Context, h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {


			if r.ContentLength > 0 {
				if data, err := ioutil.ReadAll(r.Body); err == nil {
					ctx["data"] = data
					h(w, r, p)
					return
				}
			}

			h(w, r, p)
			return
		}
	}

	return fn
}
