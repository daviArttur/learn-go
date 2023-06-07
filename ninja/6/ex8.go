package ninja6

import "fmt"

func z() func(y int) {
	return func(y int) {
		fmt.Println(y)
	}
}

func Ex8() {
	x := z()

	x(4)
}
