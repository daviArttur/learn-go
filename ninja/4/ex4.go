package ninja4

import (
	"fmt"
)

func Ex4() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)

	fmt.Println(x)

	x = append(x, []int{56, 57, 58, 60}...)

	fmt.Println(x)

}
