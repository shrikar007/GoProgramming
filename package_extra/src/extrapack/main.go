package main

import (
	"extrapack/calpack"
	"fmt"
)

func main(){
	var no,n1 int
	fmt.Println("Enter 2 to numbers")
	fmt.Scan(&no,&n1)
    sum:=calpack.Add(no,n1)
	sub:=calpack.Sub(no,n1)
	mul:=calpack.Mul(no,n1)
	div:=calpack.Div(no,n1)

	fmt.Println("Addition=",sum,"Subtraction=",sub,"multiplication=",mul,"Division=",div)

}

