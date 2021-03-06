package template

import (
	"log"
	"os"
)

// Save write the generated go file in the given filename
func (template *Template) Save(file string) {
	out, err := os.Create(file)
	if err != nil {
		log.Println("Unable to open", file)
		log.Fatal(err)
	}
	defer out.Close()

	out.WriteString("// Code generated by Gotmpl\n")
	out.WriteString("// DO NOT EDIT, I MEAN IT'S USELESS :)\n\n")
	out.WriteString("package " + template.Package + "\n")
	out.WriteString("import (\n")
	for name := range template.HiddenImports {
		out.WriteString("\t__" + name + " \"" + name + "\"\n")
	}
	for name := range template.Imports {
		out.WriteString("\t\"" + name + "\"\n")
	}
	out.WriteString(")\n")
	if template.FuncArgs == "" {
		out.WriteString("func " + template.FuncName + "(buffer *__bytes.Buffer) {\n")
	} else {
		out.WriteString("func " + template.FuncName + "(" + template.FuncArgs + ", buffer *__bytes.Buffer) {\n")
	}
	out.Write(template.Buffer.Bytes())
	out.WriteString("}\n")
}
