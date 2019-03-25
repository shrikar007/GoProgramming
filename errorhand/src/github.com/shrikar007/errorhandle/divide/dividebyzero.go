package divide

import "github.com/shrikar007/errorhandle/recoverexcep"

func Dividebyzero(num1,num2 int )(int){

	if num2==0{
		defer recoverexcep.Recovery()
		panic("Divide by zero Exception")
		return 0

	}
	return num1/num2
}
