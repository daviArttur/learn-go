package ninja2

import "fmt"

func Ex4() {

	x := 10

	fmt.Printf("%b, %d, %#x\n", x, x, x)

	d := (x << 1)

	fmt.Printf("%b, %d, %#x,", d, d, d)

}
