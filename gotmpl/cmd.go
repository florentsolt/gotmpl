package main

import (
	"flag"
	"fmt"
	"github.com/florentsolt/gotmpl/template"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	trim := flag.Bool("trim", false, "Enable trim")
	pkg := flag.String("package", "", "Specify package name")
	expr := flag.String("expr", "$", "Specify epxression token")
	tag := flag.String("tag", "go", "Specify tag name")
	flag.Parse()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fmt.Println(file.Name())
			template := template.New()
			template.Trim = *trim
			template.Expr = *expr
			template.Tag = *tag
			template.ParseFile(file.Name())
			if *pkg != "" {
				template.Package = *pkg
			}
			template.Save(file.Name() + ".go")
		}
	}
}
