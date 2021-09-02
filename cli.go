package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func run() {
	outFlag := flag.String("out", "values.yaml", "The helm values output file. By default this is values.yaml")
	flag.Parse()
	path := flag.Arg(0)

	ctx, err := newApplicationContext(path)

	if err != nil {
		log.Fatal(err)
		fmt.Print(err.Error())
		return
	}

	template := createTemplate(ctx)

	sValues, err := template.Execute()

	if err != nil {
		log.Fatal(err)
		fmt.Print(err.Error())
		return
	}

	os.WriteFile(*outFlag, []byte(sValues), 0644)
	// ctx := createAppContext(path, fileReader)

}

func createTemplate(ctx ApplicationContext) Template {
	return Template{
		Tmpl:  string((ctx).Template()),
		Funcs: funcMap(ctx),
	}
}
