package negative

import "github.com/shrikar007/errorhandle/recoverexcep"

func Negative1(num1 int)int {
	defer recoverexcep.Recovery()

	if num1<0{
		defer recoverexcep.Recovery()

		panic("Negative Number Exception")
		return 0

	}
	return num1

}
