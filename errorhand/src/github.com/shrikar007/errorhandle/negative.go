package main

import (
	"errors"
	"fmt"
	"log"
)

func Negativenu(num1,num2 int )(int,error){
	if num2==0{
		return 0,errors.New("Divided by zero Exception")
	}
	return num1/num2,nil
}
func main(){
	var n,m int
	fmt.Println("Enter numbers=")
	fmt.Scan(&n,&m)
	result, err :=Negativenu(n,m)
	 if(err!=nil){
	 	log.Fatalln(err)
	 }
	fmt.Printf("Division of %d and %d is  %d ",n,m,result)

}