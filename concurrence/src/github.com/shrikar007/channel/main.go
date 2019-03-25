package main

import "fmt"

func main(){

	ch:=make(chan bool)

	go hello(ch)
	<-ch
fmt.Println("main function")

}
func hello(chanel chan bool){

	fmt.Println("hello Shrikar")
	chanel<-true

}
