package ninja10

import (
	"fmt"
	"sync"
)

func Ex4() {
	q := make(chan int)
	c := generate(q)

	receiver(c, q)

	fmt.Println("about to exit")
}

func receiver(ch1, ch2 <-chan int) {

	for {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Printf("Channel 1 %v \n", v)
			} else {
				return
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Printf("Channel 2 %v \n", v)
			} else {
				return
			}
		}
	}
}

func generate(q chan int) <-chan int {
	c := make(chan int)
	var group sync.WaitGroup

	group.Add(2)
	go func() {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			group.Done()
		}()

		go func() {
			for i := 100; i > 0; i-- {
				q <- i
			}
			group.Done()
		}()

		group.Wait()
		close(c)
		close(q)
	}()
	return c
}
