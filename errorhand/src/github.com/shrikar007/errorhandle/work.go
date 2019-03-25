package main

import (
	"fmt"
)

func main(){

	 var input int
     fmt.Println("Enter interger only=:")
	 err,_:=fmt.Scan(&input)

	 validinput(err)
	 if input!=0{

		 fmt.Printf("%d is valid integer\n", input)
	 }
	 fmt.Println("The End")
}
func validinput(v int){

	defer recoverinput()
	if v==0{
		panic("integer required")
	}


}
func recoverinput(){
	r:=recover()
	if r==nil {
		return

	}
	fmt.Println("Recovered Runtiime Error:",r)
}
