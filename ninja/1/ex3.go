package ninja1

import (
	"fmt"
)

var a int
var b string
var c bool

func Ex3() {

	a = 42
	b = "James Bond"
	c = true

	s := fmt.Sprintf("%v\t%v\t%v\t", a, b, c)

	fmt.Println(s)
}
