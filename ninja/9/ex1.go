package ninja9

import (
	"fmt"
	"sync"
)

var group sync.WaitGroup

func Ex1() {

	group.Add(10)

	for i := 0; i < 10; i++ {
		go print(i)
	}

	group.Wait()

}

func print(id int) {

	fmt.Println(id)

	group.Done()
}
