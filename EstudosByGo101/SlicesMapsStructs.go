package main

import "fmt"

func Slices() {
	println("================================= Slices ================================")
	//------------------  SLICES ----------------------------------------
	//Diferentemente da criação de um array que é estático a criação de um slice permite o redimensionamento
	//Para auxiliar a gerenciar um slice é possível saber o seu tamanho e sua capacidade
	// len() = retorna o tamanho de um slice ou vetor
	// cap() = retorna a capacidade total de um slice
	sliceNomes := []string{"Vagner", "Simone", "Beatriz"}
	fmt.Println("Slice Original", sliceNomes, len(sliceNomes), cap(sliceNomes))

	sliceNomes = append(sliceNomes, "Juliana")
	fmt.Println("Slice com append de um nome", sliceNomes, len(sliceNomes), cap(sliceNomes))

	sliceNomes = append(sliceNomes, "Isabel", "Fábio", "Flavio", "Chiquinho", "Walter", "Claudinho")
	fmt.Println("Slice com append de múltiplos nomes", sliceNomes, len(sliceNomes), cap(sliceNomes))

	//Capacidade e tamanhos são coisas diferentes.
	//O tamanho indica o índice do último elemento
	//A capacidade indica quanto ainda pode ser alocado pelo append sem que o Go precise alocar mais memória.

	sliceNomes = append(sliceNomes, "Solange")
	fmt.Println("Ao adicionar um novo elemento o slice cresceu, porém sua capacidade não\n", sliceNomes, len(sliceNomes), cap(sliceNomes))

	//Criando um slice vazio
	sliceTamanho, sliceCapacidade := 3, 8
	sliceNovo := make([]string, sliceTamanho, sliceCapacidade)

	//Como o tamanho inicial deste slice é 3, posso usar as três primeiras posições porque estas já existem
	sliceNovo[0] = "Primeira posição"
	sliceNovo[1] = "Segunda posição"
	sliceNovo[2] = "Terceira posição"
	fmt.Println("Slice novo original", sliceNovo, len(sliceNovo), cap(sliceNovo))

	//Para utilizar as posições préviamente alocadas após o tamanho, é necessário append
	sliceNovo = append(sliceNovo, "Quarta posição")
	fmt.Println("Slice novo com 1 append", sliceNovo, len(sliceNovo), cap(sliceNovo))
	println("================================= Fim Slices ================================")
}

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

func Structs() {
	println("================================= Structs ================================")
	//Structs são tipos compostos equivalentes aos structs de C# ou Register do Delphi
	type Pessoa struct {
		Nome      string
		SobreNome string
		Idade     uint8
		Status    bool
	}

	//A atribuição pode ser campo a campo
	var vgn Pessoa
	vgn.Nome = "Vagner"
	vgn.SobreNome = "de Araujo"
	vgn.Idade = 55
	vgn.Status = true

	//A atribuição pode ser feita referenciando as propriedades e seus valores (em qualquer ordem)
	smn := Pessoa{Nome: "Simone", SobreNome: "de Araujo", Idade: 47, Status: true}

	//A atribuição pode ser feita passando apenas os valores na exata ordem das propriedades
	//Esta forma não é recomendada porque se houver uma mudança na ordem das propriedades poderemos um erro
	//ou pior um problema complicado de ser resolvido

	btrz := Pessoa{"Beatriz", "de Araujo", 7, true}

	fmt.Println(vgn)
	fmt.Println(smn)
	fmt.Println(btrz)
	println("================================= Fim Structs ================================")

}
