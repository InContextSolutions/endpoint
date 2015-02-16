package main

import (
	"fmt"
	"github.com/GolangDorks/endpoint"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func updateContext(ctx endpoint.Context, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println("...in the middleware")
		ctx["the answer"] = 42
		h(w, r, p)
	}
}

func main() {

	g := endpoint.Endpoint{
		Path:   "/foo/:bar",
		Method: "GET",
		Before: []endpoint.Middleware{updateContext},
		Control: func(ctx endpoint.Context) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				log.Println("...in the controller")
				log.Println("params:", p)
				answer, _ := ctx["the answer"]
				w.Write([]byte(fmt.Sprintf("the answer is %v\n", answer)))
			}
		},
	}

	p := endpoint.Endpoint{
		Path:   "/bar/:baz",
		Method: "POST",
		Before: []endpoint.Middleware{updateContext},
		Control: func(ctx endpoint.Context) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				log.Println("...in the controller")
				log.Println("params:", p)
				data, _ := ctx["data"]
				w.Write([]byte(fmt.Sprintf("You posted %s\n", data)))
			}
		},
	}

	router := httprouter.New()
	router.Handle(g.Method, g.Path, g.Handler())
	router.Handle(p.Method, p.Path, p.Handler())
	log.Fatal(http.ListenAndServe(":8080", router))
}
