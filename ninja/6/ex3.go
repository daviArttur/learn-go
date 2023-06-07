package ninja6

import "fmt"

func Ex3() {
	defer fmt.Println("LOG: Deveria ser o primeiro! Porém foi o último por usar o 'defer'")
	fmt.Println("LOG: Deveria ser o último")
}
