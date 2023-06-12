package ninja10

import "fmt"

func Ex6() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
