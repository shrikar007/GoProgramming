package main

import "fmt"
var num1,num2 int
func main() {

	fmt.Println("Enter 2 numbers:")
	fmt.Scan(&num1,&num2)
	 add(num1,num2)

}

func add(no1,no2 int){

	addition:=no1+no2

	fmt.Println("addition=", addition)


}
