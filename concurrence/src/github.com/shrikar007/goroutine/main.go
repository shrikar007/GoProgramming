package main

import (
	"fmt"
	"time"
)

func main(){
	go iterate1()
	go iterate2()
	time.Sleep(2*time.Second)

	fmt.Println("The End")

}
func iterate1(){

	for i:=1;i<=10;i++{
		time.Sleep(100)

		fmt.Println(i)
	}
}
func iterate2(){

	for i:=11;i<=20;i++{
		time.Sleep(100)

		fmt.Println(i)
	}
}
