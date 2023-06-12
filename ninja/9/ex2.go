package ninja9

import "fmt"

type Person struct {
	Name string
}

type Human interface {
	speak()
}

func Ex2() {
	mario := Person{
		Name: "Mario",
	}

	sayAnything(&mario)
}

func (p *Person) speak() {
	fmt.Println("Its me", (*p).Name)
}

func sayAnything(h Human) {
	h.speak()
}
