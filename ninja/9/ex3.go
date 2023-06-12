package ninja9

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int64

func Ex3() {

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			value := counter
			runtime.Gosched()
			value++
			counter = value
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
