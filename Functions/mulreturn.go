package main

import "fmt"

func main(){
	var x,y int
	fmt.Println("Enter 2 values=")
	fmt.Scan(&x,&y)
	resultsum,resultsub:=addsub(x,y)
	fmt.Println("Addition=",resultsum,"Substraction=",resultsub)
}

func addsub(a,b int)(sum int, sub int){
	sum=a+b
	sub=a-b

	return

}
