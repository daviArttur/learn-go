package ninja6

import "fmt"

func Ex2(x []int) {
	total := 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}
