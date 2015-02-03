package main

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
}
`

var endpointTestTemplate = `package {{.Package}}

import (
` + "\t" + `"github.com/stretchr/testify/assert"
` + "\t" + `"net/http"
` + "\t" + `"net/http/httptest"
` + "\t" + `"testing"
)

func run{{.Name}}() *httptest.ResponseRecorder {
` + "\t" + `r, _ := http.NewRequest("{{.Method}}", "http://example.com/foo", nil)
` + "\t" + `w := httptest.NewRecorder()
` + "\t" + `// modify request here
` + "\t" + `{{.Name}}.Handler().ServeHTTP(w, r)
` + "\t" + `return w
}

func Test{{.Name}}(t *testing.T) {
` + "\t" + `w := run{{.Name}}()
` + "\t" + `assert.Equal(t, w.Code, -1, "status code check not implemented")
` + "\t" + `t.Error("test not implemented")
}

func Benchmark{{.Name}}(b *testing.B) {
` + "\t" + `for n := 0; n < b.N; n++ {
` + "\t\t" + `run{{.Name}}()
` + "\t" + `}
}
`
