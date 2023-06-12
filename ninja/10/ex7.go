package ninja10

import (
	"fmt"
	"sync"
)

func Ex7() {

	ch := make(chan string)

	go createProducers(10, ch)

	for a := range ch {
		fmt.Println(a)
	}
}

func createProducers(q int, ch chan string) {

	var group sync.WaitGroup

	group.Add(q)

	for i := 0; i < q; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {

				ch <- fmt.Sprintf("By go func: %v value: %v", i, j)
			}
			group.Done()
		}(i)
	}
	group.Wait()
	close(ch)
}
