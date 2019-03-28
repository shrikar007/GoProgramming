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
func main(){
	t1, err := template.ParseFiles("login.html")
	//t2, err1 := template.ParseFiles("display1.html")
	if err != nil{
		log.Fatalln(err)
	}
	h1:=func (w http.ResponseWriter,r *http.Request){

		var l Login
		l.username = r.FormValue("user")
		l.password = r.FormValue("pass")
		t1.Execute(w,l)
		if l.password =="shri"{
			fmt.Println("Login successfull")
		} else {
			fmt.Println("Login error")
		}
		//fmt.Println( l.username)
		//fmt.Println( l.password)
		if err != nil{
			log.Fatal(err)
		}
		//t2.Execute(w,l)
	}
	http.HandleFunc("/login",h1)
	fmt.Println(http.ListenAndServe(":5004",nil))
}

