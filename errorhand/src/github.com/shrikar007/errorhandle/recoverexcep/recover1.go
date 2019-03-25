package recoverexcep

import "fmt"

func Recovery(){

	r:=recover()
	if r!=nil {

		fmt.Println("Recovered:", r)
	}
}

