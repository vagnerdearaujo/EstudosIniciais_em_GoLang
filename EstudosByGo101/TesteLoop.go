package main

import "fmt"

func TesteLoop() {
	nomes := []string{"Vagner", "Simone", "Beatriz", "Juliana", "Isabel", "Fábio", "Flavio", "Chiquinho", "Walter", "Claudinho"}

	//For padrão
	for idx := 0; idx < len(nomes); idx++ {
		fmt.Println(nomes[idx])
	}

	//Simular ForEach com captura do índice
	for idx, nome := range nomes {
		fmt.Println("Índice:", idx, "=", nome)
	}

	//Simular ForEach desprezando o índice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//Simular um do while
	var idx int //A simples criação de uma variável já atribui o valor padrão que neste caso será 0
	for idx < len(nomes) {
		fmt.Println(nomes[idx])
		idx++
	}

	/*
		É possível criar um loop infinito apenas colocando uma clausula for
		for {
			comandos ...
		}
	*/
}
