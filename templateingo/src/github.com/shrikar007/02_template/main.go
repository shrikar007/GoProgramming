package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Fname        string
	Lname         string
}

func main() {
	http.HandleFunc("/fullname", HomePage)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

		HomePageVars := PageVariables{
			"Shrikar",
			"Vaitala",
	}

	t, err := template.ParseFiles("fullname.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

