package main

import (

	"text/template"
	"log"
	"os"
)
var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseGlob("temp/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "World!")
	if err != nil {
		log.Fatal(err)
	}
}