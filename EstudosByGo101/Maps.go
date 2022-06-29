package main

import "fmt"

func Maps() {
	println("================================= Maps ================================")
	//Map é uma coleção do tipo Chave / Valor
	//As chaves não precisam ter o mesmo tipo dos valores

	//Criação de um MAP
	idades := make(map[string]uint8)

	//Adicionando ítens ao MAP
	idades["Vagner"] = 55
	idades["Simone"] = 47
	idades["Beatriz"] = 7
	idades["Juliana"] = 21

	//Para percorrer todo o mapa
	for key := range idades {
		fmt.Println(key, idades[key])
	}

	//Se uma chave não existente for passada para o map o retorno será o valor básico, neste caso 0

	//Não é possível passar o map como referência, ou seja:
	//&idade["Vagner"] geraria um erro de compilação.

	//Como saber se um determinada chave existe?
	idade, existe := idades["Vagner"]
	fmt.Println("Quando a chave existe o valor da variável \"existe\" será true: idade =", idade, "; existe =", existe)

	idade, existe = idades["Chiquinho"]
	fmt.Println("Quando a chave não existe o valor da variável \"existe\" será false: idade =", idade, "; existe =", existe)
	println("================================= Fim Maps ================================")
}
