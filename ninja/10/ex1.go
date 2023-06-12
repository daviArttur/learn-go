package ninja10

import (
	"fmt"
)

func Ex1() {
	case1()
	case2()
}

func case1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func case2() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
