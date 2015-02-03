package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var (
	pkg    = flag.String("package", "", "name of the package")
	point  = flag.String("endpoint", "", "name of the endpoint")
	path   = flag.String("path", "", "url path")
	method = flag.Bool("post", false, "use POST (default is GET)")
	output = flag.String("output", "", "output directory (default is current directory")
)

func main() {
	flag.Parse()

	if *pkg == "" || *point == "" || *path == "" {
		panic("must set package, endpoint, and path")
	}

	meth := "GET"
	if *method {
		meth = "POST"
	}

	stub := template.Must(template.New("T").Parse(endpointTemplate))
	stubtest := template.Must(template.New("T").Parse(endpointTestTemplate))

	data := struct {
		Package string
		Name    string
		Path    string
		Method  string
	}{
		Package: *pkg,
		Name:    *point,
		Path:    *path,
		Method:  meth,
	}

	var dir string
	var err error
	if *output == "" {
		dir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		dir += "/"
	} else {
		dir = *output
	}
	fname := strings.ToLower(*point)

	buf := bytes.Buffer{}
	stub.Execute(&buf, data)
	ioutil.WriteFile(dir+fname+".go", buf.Bytes(), 0644)

	buf.Reset()
	stubtest.Execute(&buf, data)
	ioutil.WriteFile(dir+fname+"_test.go", buf.Bytes(), 0644)
}
