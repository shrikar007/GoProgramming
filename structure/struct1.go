package main

import "fmt"

type student struct{

	fname string
	lname string
	marks int

}
func main(){
	s:=student{"shrikar","vaitala",50}
	

	fmt.Println(fullname(s))

}
func fullname(s1 student) string{

	return s1.fname + " "+s1.lname


}
