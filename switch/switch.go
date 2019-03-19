package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	switch name {
	case "Shrikar":
		fmt.Println("Hello Shrikar")
	case "Mohanish":
		fmt.Println("Hello, Mohanish!")
		  break
              
	case "Rohit":
		fmt.Println("Hello, Rohit!")
                fallthrough
   
	case "Ankit":
		fmt.Println("Hello, Ankit!")
	case "Soni":
		fmt.Println("Hello, Soni!")
		fallthrough
	case "Dinesh":
		fmt.Println("Hello Dinesh")
	default:
		fmt.Println("Hello, " + name + " you are not a part of GoLang Training. Would you like to join us?")
	}
}
