package ninja10

import (
	"fmt"
	"time"
)

func Ex3() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}

}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		time.Sleep(1 * time.Second)
		c <- 999
		close(c)
	}()

	return c
}
