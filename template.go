package main

import (
	"strings"
	"text/template"
)

type Template struct {
	Tmpl  string
	Funcs map[string]interface{}
}

func (tmpl *Template) Execute() (string, error) {

	var buffer strings.Builder
	err := template.Must(
		template.New("filename").
			Funcs(tmpl.Funcs).
			Parse(tmpl.Tmpl)).
		Execute(&buffer, new(interface{}))

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
