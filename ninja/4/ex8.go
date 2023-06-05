package ninja4

import "fmt"

func Ex8() {

	x := map[string]string{
		"silva_cafezinis":  "fazer café",
		"joelson_queixada": "boxe",
	}

	for _, person := range x {
		for _, personData := range person {
			fmt.Println(personData)
		}
	}
}
