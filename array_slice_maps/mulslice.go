package main

import "fmt"

func main() {
	var scored [][]string

	marks := make([]string, 3)
	marks[0] = "50"
	marks[1] = "60"
	marks[2] = "70"

	scored = append(scored, marks)

	students := make([]string, 3)
	students[0] = "shrikar"
	students[1] = "dinesh"
	students[2] = "mohanish"

	 scored = append(scored, students)

	fmt.Println(scored)

}
