package main

import (
	"bytes"
	"flag"
	"fmt"
	"text/template"
)

var (
	pkg    = flag.String("package", "", "name of the package")
	point  = flag.String("endpoint", "", "name of the endpoint")
	path   = flag.String("path", "", "url path")
	method = flag.Bool("post", false, "use POST (default is GET)")
)

var endpointTemplate = `package {{.Package}}

import (
` + "\t" + `"github.com/GolangDorks/endpoint"
` + "\t" + `"net/http"
)

// {{.Name}} accepts {{.Method}} requests.
var {{.Name}} = endpoint.Endpoint{
` + "\t" + `Path:   "{{.Path}}",
` + "\t" + `Method: endpoint.{{.Method}},
` + "\t" + `Before: []endpoint.Middleware{},
` + "\t" + `Control: func(ctx endpoint.Context) http.Handler {
` + "\t\t" + `return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
` + "\t\t\t" + `// this is the controller body
` + "\t\t\t" + `w.WriteHeader(http.StatusNotImplemented)
` + "\t\t" + `})
` + "\t" + `},
}`

func main() {
	flag.Parse()

	if *pkg == "" || *point == "" || *path == "" {
		panic("must set package, endpoint, and path")
	}

	meth := "GET"
	if *method {
		meth = "POST"
	}

	tmpl := template.Must(template.New("T").Parse(endpointTemplate))

	buf := bytes.Buffer{}
	tmpl.Execute(&buf, struct {
		Package string
		Name    string
		Path    string
		Method  string
	}{
		Package: *pkg,
		Name:    *point,
		Path:    *path,
		Method:  meth,
	})

	fmt.Println(buf.String())
}
