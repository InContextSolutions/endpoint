package endpoint

import "net/http"

type handler struct {
	before  []Middleware
	control Controller
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := make(Context)

	final := h.control(ctx)
	for i := len(h.before) - 1; i >= 0; i-- {
		final = h.before[i](ctx, final)
	}

	final.ServeHTTP(w, r)
}
