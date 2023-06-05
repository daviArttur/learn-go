package ninja4

import "fmt"

func Ex7() {

	x := [][]string{
		{"Name1", "Surname1", "Hobby1"},
		{"Name2", "Surname2", "Hobby2"},
		{"Name3", "Surname3", "Hobby3"},
	}

	for _, person := range x {
		for _, personData := range person {
			fmt.Println(personData)
		}
	}
}
