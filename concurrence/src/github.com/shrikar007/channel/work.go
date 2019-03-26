package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	fmt.Println("hello")
}

func main() {
	sendch := make(chan int)
	go sendData(sendch)
	val,flag:= <-sendch
	fmt.Println(val)
	fmt.Println(flag)
	fmt.Println("main")
}
