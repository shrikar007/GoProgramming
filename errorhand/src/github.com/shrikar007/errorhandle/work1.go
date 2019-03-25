package main

import "fmt"

func main(){

	fmt.Println("Enter numbers=")
	var n,m int
	fmt.Scan(&n,&m)
	a:=func(num1, num2 int) {

		sum:=num1+num2
		fmt.Println("addition=",sum)
		}


	a(n,m)
}
