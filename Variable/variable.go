package main

import "fmt"

      var num int //basic declaration
	


	// multiple declarations
	var hello, world, message string 
	var a, b, c int
func main() {

        var num1 = 10 //inferred type
	num2:=45 // shorthand declaration
	num=20
	hello = "Hello,"	
	world = "World!"	
	a = 1
        fmt.Printf("num(type) - %T (value) - %v\n", num,num)
        fmt.Printf("num1(type) - %T (value) - %v\n", num1,num1)
        fmt.Printf("num2(type) - %T (value) - %v\n", num2,num2)
	fmt.Printf("hello(type) - %T (value) - %v\n", hello, hello)
	fmt.Printf("world(type) - %T (value) - %v\n", world, world)
	fmt.Printf("message(type) - %T (value) - %v\n", message, message)
	fmt.Printf("a(type) - %T (value) - %v\n", a, a)
	fmt.Printf("b(type) - %T (value) - %v\n", b, b)
	fmt.Printf("c(type) - %T (value) - %v", c, c)

}
