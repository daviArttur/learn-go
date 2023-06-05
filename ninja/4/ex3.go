package ninja4

import (
	"fmt"
)

func Ex3() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers[:3])
	fmt.Println(numbers[4:])
	fmt.Println(numbers[1 : len(numbers)-1])
}
