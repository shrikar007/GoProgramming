package main

import (
	"log"
	"net/http"
	"text/template"
)
var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseGlob("temp/*"))
}

func main() {
	http.HandleFunc("/slice",HomePage)
	log.Fatal(http.ListenAndServe(":8086",nil))



}
func HomePage(w http.ResponseWriter, r *http.Request) {

	s := []string{"shrikar", "dinesh", "mohanish", "prathamesh"}
	err := tpl.ExecuteTemplate(w, "slice.html", s)
	if err != nil {
		log.Fatal(err)
	}
}