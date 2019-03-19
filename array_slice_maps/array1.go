package main

import "fmt"

func main() {
	var x [5]int

	fmt.Printf("Type: %T\n", x)
	fmt.Println("Enter array values")
        for i := 0; i < 5; i++ {
            fmt.Scan(&x[i])
        }
	fmt.Println("Length of Array", len(x))
	fmt.Println("Value at 2th position", x[1])
	x[1] = 27
	fmt.Println("Value at 2th position", x[1])
}
