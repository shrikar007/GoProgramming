package main

import (
	"fmt"
	"github.com/shrikar007/math"
)

func main(){
	var n2 int
	var n1 int
	fmt.Println("Enter the number=")
	fmt.Scan(&n1,&n2)
	sum := math.Add(n2,n1)
	fmt.Println("Addition =",sum)

}