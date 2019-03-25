package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fname, err := os.Open("test.txt")

	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		return

	}
	fmt.Println(fname.Name(), "file open successfull")
}
