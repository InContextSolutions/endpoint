package main

import (
	"fmt"
	"github.com/GolangDorks/endpoint"
	"log"
	"net/http"
)

func updateContext(ctx endpoint.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("...in the middleware")
		ctx["the answer"] = 42
		h.ServeHTTP(w, r)
	})
}

func main() {

	g := endpoint.Endpoint{
		Path:   "/foo",
		Method: endpoint.GET,
		Before: []endpoint.Middleware{updateContext},
		Control: func(ctx endpoint.Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					log.Println("...in the controller")
					answer, _ := ctx["the answer"]
					w.Write([]byte(fmt.Sprintf(
						"the middleware told me the answer is %v", answer)))
				})
		},
	}

	p := endpoint.Endpoint{
		Path:   "/bar",
		Method: endpoint.POST,
		Before: []endpoint.Middleware{updateContext},
		Control: func(ctx endpoint.Context) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					log.Println("...in the controller")
					data, _ := ctx["data"]
					w.Write([]byte(fmt.Sprintf(
						"You posted %v", data)))
				})
		},
	}

	http.Handle(g.Path, g.Handler())
	http.Handle(p.Path, p.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
