package ninja4

import (
	"fmt"
)

func Ex2() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, currentNumber := range numbers {
		fmt.Printf("%T \n", currentNumber)
	}
}
