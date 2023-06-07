package ninja7

import "fmt"

type Person struct {
	name string
}

func changeMe(p *Person) {
	(*p).name = "Robiss"
}

func Ex2() {
	person1 := Person{
		name: "John",
	}

	fmt.Println(person1)

	changeMe(&person1)

	fmt.Println(person1)
}
