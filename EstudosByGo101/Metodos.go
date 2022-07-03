package main

import "fmt"

type Pessoa struct {
	Nome      string
	SobreNome string
	Idade     uint8
	Status    bool
	cpf       string
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

//String:Este método permite que uma saída formatada seja apresentada quando a struct for impressa "seca"
func (p Pessoa) String() string {
	var ativo string
	if p.Status {
		ativo = "Ativo"
	} else {
		ativo = "Inativo"
	}
	return fmt.Sprintf("Nome: %s %s - CPF: %s (Status: %s)", p.Nome, p.SobreNome, p.cpf, ativo)
}

func (c Categoria) Print() string {
	/*
		var result string
		if c.Pai == nil {
			result = fmt.Sprint(c.Nome)
		} else {
			result = fmt.Sprintf("%s (%s)", c.Nome, c.Pai.Nome)
		}
		return result
	*/

	return fmt.Sprint(c.Nome)
}

func Metodos() {
	//Se houver a necessidade de aninhar a categoria dentro dela é obrigatório o uso de * (ponteiro)

	p := Pessoa{
		Nome:      "Vagner",
		SobreNome: "de Araujo",
		Idade:     55,
		Status:    true,
		cpf:       "123.456.789-00",
	}

	fmt.Println(p.Nome, p.SobreNome, p.Idade, p.cpf)

	cat := Categoria{Nome: "Categoria-A", Pai: &Categoria{Nome: "Pai da Categoria"}}

	if !cat.HasParent() {
		fmt.Println("A categoria", cat.Nome, "\"Has no parent\"")
	} else {
		fmt.Println("A categoria", cat.Nome, "é filha de", cat.Pai.Nome)
	}

	fmt.Println(p)
	fmt.Println(cat)
}
