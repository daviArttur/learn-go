package ninja4

import "fmt"

func Ex6() {

	estados := make([]string, 0, 26)
	fmt.Print(len(estados), cap(estados), "\n")

	estados = []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
		"Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
		"Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
	}

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
