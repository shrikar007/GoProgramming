package main

import (
	"fmt"
	"time"
)


func main(){
	ch1:=make(chan bool ,2)
	//ch2:=make(chan bool)
	go iterat1(ch1)
	//<-ch1
	go iterat2(ch1)
	//time.Sleep(1*time.Second)
    <-ch1
	<-ch1

	fmt.Println("\nThe End")

}



func iterat1(ch chan bool){



	for  i:=1; i<=10;i++{

		time.Sleep(100)



		fmt.Print(" ",i)
	}
	ch<-true
}
func iterat2(ch chan bool){

	for i:=11;i<=20;i++{
		time.Sleep(100)

		fmt.Print(" ",i)
	}
	ch<-true
}
