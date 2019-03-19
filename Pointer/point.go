package main

import "fmt"

func main() {
	a := 24
	b := &a 

	fmt.Println(a, *b) 

	c := *b 
	a = 30 

	fmt.Println(a, *b, c) 
}
