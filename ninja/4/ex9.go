package ninja4

import "fmt"

func Ex9() {

	x := map[string]string{
		"silva_cafezinis":  "fazer café",
		"joelson_queixada": "boxe",
	}

	x["nobre_cebola"] = "odeia cebola"
	for _, person := range x {
		fmt.Println(person)
	}
}
