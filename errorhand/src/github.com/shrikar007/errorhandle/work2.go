package main

import (
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(o order)  {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	j := fmt.Sprintf("Select * from order \n%d, %d", o.ordId, o.customerId)
	//fmt.Printf("%T\n",i)
	fmt.Println(i)
	fmt.Println(j)


}

func main() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	createQuery(o)
}