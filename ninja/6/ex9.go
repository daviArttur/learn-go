package ninja6

import "fmt"

func Ex9() {
	x := func(callback func()) {
		callback()
	}

	x(func() {
		fmt.Println("Callback here")
	})
}
