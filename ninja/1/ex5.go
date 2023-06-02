package ninja1

import "fmt"

type Number int

var myInt int
var myNumber Number

func Ex5() {

	fmt.Println(number)
	fmt.Printf("%T\n", number)
	number = 42
	fmt.Println(number)

	myNumber = Number(myInt)

	fmt.Println(myNumber)
	fmt.Printf("%T", myNumber)
}
