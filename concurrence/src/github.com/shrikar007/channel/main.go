package main

import (
	"fmt"
	"time"
)

func main(){

	ch:=make(chan bool)
	fmt.Println("main before entering into hello")

	go hello(ch)
	<-ch
	time.Sleep(1*time.Second)
fmt.Println("main function")

}
func hello(chanel chan bool){

	fmt.Println("in hello function wait for 4 seconds")
	//time.Sleep(4*time.Second)

	fmt.Println("after 4 seconds return back to main")


	//chanel<-true

}
