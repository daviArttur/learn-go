package ninja6

import "fmt"

func closure() func() {
	calls := 0
	return func() {
		calls++
		fmt.Println(calls)
	}
}

func Ex10() {
	x := closure()
	x()
	x()
	x()

	x2 := closure()
	x2()
	x2()
	x2()
}
