package ninja6

import "fmt"

type Person struct {
	Name    string
	Surname string
}

func Ex4() {
	person1 := Person{
		Name:    "John",
		Surname: "Doe",
	}
	person1.SayHello()
}

func (p Person) SayHello() {
	fmt.Println(p.Name + " " + p.Surname + " say hello!")
}
