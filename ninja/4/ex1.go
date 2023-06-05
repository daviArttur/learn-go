package ninja4

import (
	"fmt"
)

func Ex1() {

	numbers := [5]int{1, 2, 3, 4, 5}

	for _, currentNumber := range numbers {
		fmt.Printf("%T %v \n", currentNumber, currentNumber)
	}
}
