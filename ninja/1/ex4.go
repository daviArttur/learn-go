package ninja1

import "fmt"

var number int

func Ex4() {

	fmt.Println(number)
	fmt.Printf("%T\n", number)
	number = 42
	fmt.Println(number)
}
