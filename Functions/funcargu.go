package main

import "fmt"

func main(){
	var nam string

	fmt.Println("Enter your name ")
	fmt.Scan(&nam)
	myname(nam)

}
func myname(name string){
	fmt.Println("Your name is",name)
}

