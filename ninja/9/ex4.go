package ninja9

import (
	"fmt"
	"runtime"
	"sync"
)

func Ex4() {

	type MutexCounter struct {
		mutex sync.Mutex
		value int
	}

	counter := MutexCounter{
		mutex: sync.Mutex{},
		value: 0,
	}

	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			counter.mutex.Lock()
			counter.value++
			counter.mutex.Unlock()
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.value)
}
