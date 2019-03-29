package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
type Login struct {
	username  string
	password string
}
var pass= "shri"
func main(){
	http.HandleFunc("/login1",h1)
	fmt.Println(http.ListenAndServe(":4505",nil))
}
func h1 (w http.ResponseWriter,r *http.Request) {
	var l Login
	var Success bool
	t1, err := template.ParseFiles("login.html")

	if err != nil {
		log.Fatal(err)
	}
	l.username = r.FormValue("user")
	l.password = r.FormValue("pass")
	if r.FormValue("pass") == pass {
		Success=true
	} else  {
		Success=false
	}
	err1 := t1.ExecuteTemplate(w,"login.html",Success)
	if err1!=nil{
		log.Fatal(err1)
	}



}
