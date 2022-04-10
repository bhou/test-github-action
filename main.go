package main

import (
	"fmt"
	"strings"
	"text/template"
)

type Context struct {
	Os   string
	Arch string
	Exe  string
}

func main() {
	fmt.Println("Hello World")

	ctx := Context{
		Os:   "windows",
		Arch: "arm",
		Exe:  "",
	}

	t, err := template.New("test").Parse("{{.Os}}/{{.Arch}}/src/test{{if eq .Os \"windows\"}}.bat{{else if eq .Os \"linux\"}}.sh{{end}}")
	if err != nil {
		fmt.Println(err.Error())
	}

	builder := strings.Builder{}
	err = t.Execute(&builder, ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(builder.String())
}
