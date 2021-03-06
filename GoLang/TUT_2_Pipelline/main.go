package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, 42) //Printout "42" in the html file
	if err != nil {
		log.Fatalln(err)
	}
}

// Read this
// https://en.wikipedia.org/wiki/Pipeline_(computing)
