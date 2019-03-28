package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
var tpl *template.Template
func init(){
	//fmt.Println(template.ParseGlob("display.html"))
	tpl= template.Must(template.ParseGlob("display.html"))
}

func main(){
	h1:=func (w http.ResponseWriter,r *http.Request){
		s1:=map[int]string{50:"shrikar",14:"dinesh",59:"mohanish",89:"prathamesh"}
         err:=tpl.ExecuteTemplate(w,"display.html",s1)
        //fmt.Println(tpl.ExecuteTemplate(w,"display.html",s1))
         if err != nil{
         	log.Fatal(err)
		 }
	}
//fmt.Println(h1)
	http.HandleFunc("/display",h1)
	fmt.Println(http.ListenAndServe(":4008",nil))
}
