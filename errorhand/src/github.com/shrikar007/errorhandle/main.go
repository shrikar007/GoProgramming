package main

import (
	"fmt"
	"github.com/shrikar007/errorhandle/divide"
	"github.com/shrikar007/errorhandle/negative"
)
func main(){
	var n,m,x ,ch int
	fmt.Println("1: Divide:\n2:negative:")
	fmt.Println("Enter your Choice :")
	fmt.Scan(&ch)
	switch ch {
	case 1:fmt.Println("Enter numbers=")
		                  fmt.Scan(&n,&m)
							result :=divide.Dividebyzero(n,m)
							if result!=0 {
								fmt.Printf("Division of %d and %d is  %d \n",n,m,result)}
	case 2:fmt.Println("Enter numbers=")
		                        fmt.Scan(&x)
							result :=negative.Negative1(x)
							if result!=0 {
								fmt.Printf("%d is valid Number\n",x)}
	default:
		fmt.Println("Invalid Choice!!!!!!!!!!")
	}

	fmt.Println("The End")


}


