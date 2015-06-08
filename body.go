package endpoint

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func readBody() Middleware {
	fn := func(ctx Context, h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {



			h(w, r, p)
			return
		}
	}

	return fn
}
