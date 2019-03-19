package main

import "fmt"

func main(){
	var a int
	fmt.Println("Enter the Number=")
	fmt.Scan(&a)
	sqresult:=square(a)
	cuberesult:=cube(a)
	fmt.Println("the Square is=",sqresult)
	fmt.Println("the Cube is=",cuberesult)
}
func square(x int )int {

	return x*x
}
func cube( x int)int{

	return x*x*x
}