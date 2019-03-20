package mypkg

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Println("Hi, My name is", p.Name)
}

type Student struct {
	Person // promotion

	School string
}

func (s *Student) Introduce() {
	s.Person.Introduce()
	s.StudiesAt()
}

func (s *Student) StudiesAt() {
	fmt.Println("I study at ", s.School)
}