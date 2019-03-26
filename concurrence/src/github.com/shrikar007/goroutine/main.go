package main

import (
	"fmt"
	"time"
)


func main(){


	 go iterate1()


	go iterate2()
	time.Sleep(11*time.Second)

	fmt.Println("\nThe End")

}



func iterate1(){



	for  i:=1; i<=10;i++{

		time.Sleep(1*time.Second)




		fmt.Print(" ",i)
	}

}
func iterate2(){

	for i:=11;i<=20;i++{
		time.Sleep(1*time.Second)

		fmt.Print(" ",i)
	}

}
