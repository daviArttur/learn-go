package ninja10

import (
	"fmt"
)

func Ex2() {
	cs := make(chan int)

	go func(chan<- int) {
		cs <- 42
	}(cs)

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
