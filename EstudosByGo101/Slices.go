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
