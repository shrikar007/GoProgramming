package main

import (
	"fmt"
	"github.com/shrikar007/package-demo/pkg"
)

func main() {
	//fmt.Println("abcd")
	//mypkg.PrintMe("something")
	//
	//sum := math.Add(5, 5)
	//fmt.Println("sum", sum)

	//person := mypkg.Person{
	//	Name: "shrikar",
	//}
	//
	//fmt.Printf("%T\n", person)
	//
	//person.Introduce()
	//
	//student := mypkg.Student{}
	//
	//fmt.Printf("%T\n", student)
	//
	//student.Name = "shrikar"
	//student.School = "mbbp"
	//
	//student.Introduce()
	//student.StudiesAt()
	//
	//student.Person.Introduce()

	square := pkg.Square{}
	square.Side = 10.5

	PrintAreaAndCircumference(&square)
	//fmt.Println(square,"\n")

	rectangle := pkg.Rectangle{}
	rectangle.Length = 10.5
	rectangle.Width = 5.5

	PrintAreaAndCircumference(&rectangle)

	circle := pkg.Circle{}
	circle.Radius = 2.5

	PrintAreaAndCircumference(&circle)
}

func PrintAreaAndCircumference(shape pkg.Shape) {
	fmt.Println("area: ", shape.Area(), "circumference: ", shape.Circumference())
}