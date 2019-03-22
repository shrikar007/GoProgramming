package main

import "fmt"

type student struct{

	fname string
	lname string
}
func (s student) fullname(){
    fmt.Println("First name=",s.fname,"Last name=",s.lname)

}
func main(){

	s1:=student{
		"Shrikar",
		"vaitala",
	}
	s2:=student{
		"Dinesh",
		"Jangam",
	}
	s3:=student{
		"moni",
		"sargam",
	}

	 s1.fullname()
	defer s2.fullname()
	s3.fullname()
	fmt.Println("Hello ")

}

