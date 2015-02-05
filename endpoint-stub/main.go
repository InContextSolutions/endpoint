package main

import (
	"bytes"
	"flag"
	"fmt"
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
	force  = flag.Bool("force", false, "force overwrite of existing files")
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

	fullpath := dir + fname + ".go"
	buf := bytes.Buffer{}
	if _, err := os.Stat(fullpath); err != nil || *force == true {
		stub.Execute(&buf, data)
		ioutil.WriteFile(fullpath, buf.Bytes(), 0644)
	} else {
		fmt.Printf("'%s' exists; skipping\n", fullpath)
	}

	fullpath = dir + fname + "_test.go"
	if _, err := os.Stat(fullpath); err != nil || *force == true {
		buf.Reset()
		stubtest.Execute(&buf, data)
		ioutil.WriteFile(fullpath, buf.Bytes(), 0644)
	} else {
		fmt.Printf("'%s' exists; skipping\n", fullpath)
	}
}
