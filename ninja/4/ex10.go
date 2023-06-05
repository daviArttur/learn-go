package ninja4

import "fmt"

func Ex10() {

	x := map[string]string{
		"silva_cafezinis":  "fazer café",
		"joelson_queixada": "boxe",
	}

	delete(x, "joelson_queixada")
	for _, person := range x {
		fmt.Println(person)
	}
}
