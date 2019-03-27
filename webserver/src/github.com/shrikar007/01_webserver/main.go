package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/firstweb",handle)
   err:= http.ListenAndServe(":2000",nil)

   if err != nil{
   	  log.Fatal(err)
   }

}

func handle(w http.ResponseWriter,r *http.Request){


	fmt.Fprint(w,"hello world")
}