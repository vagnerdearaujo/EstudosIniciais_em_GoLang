package main

import "fmt"

func Structs() {
	type Pessoa struct {
		Nome      string
		SobreNome string
		Idade     uint8
		Status    bool
	}

	println("================================= Structs ================================")
	//Structs são tipos compostos equivalentes aos structs de C# ou Register do Delphi

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
