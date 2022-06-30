package main

import "fmt"

func Metodos() {
	type Pessoa struct {
		Nome      string
		SobreNome string
		Idade     uint8
		Status    bool
		cpf       string
	}

	//Se houver a necessidade de aninhar a categoria dentro dela é obrigatório o uso de * (ponteiro)
	type Categoria struct {
		Nome string
		Pai  *Categoria
	}

	p := Pessoa{
		Nome:      "Vagner",
		SobreNome: "de Araujo",
		Idade:     55,
		Status:    false,
		cpf:       "123.456.789-00",
	}

	fmt.Println(p.Nome, p.SobreNome, p.Idade, p.cpf)

}
