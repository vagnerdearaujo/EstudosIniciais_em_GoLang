package main

import "fmt"

func main() {
	fmt.Printf("A soma de X + Y = %v\n", soma(25, 74))
	fmt.Printf("O resultado de somadiversos = %v\n", somadiversos(25, 14, 13, 67, 9, 10, 8, 7, 55, 47))

	//Cria uma "instância" de carro e atribui os valores
	car := carro{Nome: "Yaris XL", Marca: "Toyota", Modelo: "2022/2023"}

	//Chama o método odometro
	car.odometro(1910)

	//O método odometro muda a cor para Prata
	fmt.Println("A cor do carro é:" + car.Cor)

	var car2 carro
	car2.Cor = "Azul"

	fmt.Println("Cor do carro 2: " + car2.Cor)

	//Como são instâncias separadas, mesmo chamando o método odometro para car a cor do car2 não é alterada de azul
	//para prata
	car.Nome = "Yaris Lindo"
	car.odometro(1910)

	fmt.Println("Cor do carro 2: " + car2.Cor)

	//Outra forma de inicializar as propriedades.
	car2 = carro{"City", "Honda", "EXL 2012/2013", "Cinza"}
	fmt.Printf("O carro %v da marca %v modelo %v e cor %v já foi vendido\n", car2.Nome, car2.Marca, car2.Modelo, car2.Cor)

	//Também é possível criar um método que inicialize as propriedades
	car2.inicializa("Face", "Chery", "2012/2013", "Branco")
	fmt.Println("O carro", car2.Nome, "da marca", car2.Marca, "modelo", car2.Modelo, "e cor", car2.Cor, "também já foi vendido")

}

//soma: esta função tem um retorno nomeado, nomear o resultado não é uma obrigação, porém evita de ter que criar
// uma variável para tarefas que exijam a execução de cálculos ou considerações.
func soma(x, y int) (result int) {
	result = x + y
	return result
}

//somadiversos: ao declarar um parâmetro seguido de "..." e o tipo é como se tivéssemos criados uma lista de parâmetros
func somadiversos(x ...int) (result int) {
	result = 0
	//o for exige uma variável índice ao colocar um "_" indicamos que nenhuma variável será passada
	// o range x explora a "lista x" e coloca item a item na variável "valor"
	for _, valor := range x {
		result += valor
	}
	return
}

//O GO Lang não implementa Orientação à Objetos de forma tradicional, portanto, não existe a declaração de uma classe
//o que existe é a declaração de um tipo definido pelo usuário (struct) que terá propriedades.
// os métodos de uma "classe" em GO Lang são funções com um tipo especial de assinatura que "atacha" o método à estrutura
type carro struct {
	Nome   string
	Marca  string
	Modelo string
	Cor    string
}

//Declara o método anda como se fosse um método da "classe" carro que na verdade é uma estrutura
func (c *carro) odometro(distancia int) {
	//Como a variável c recebe um ponteiro para a instância, é possível modificar a propriedade Cor.
	fmt.Printf("O carro %v andou %v kms\n", c.Nome, distancia)
	c.Cor = "Prata"
}

//Declara um método equivalente ao create das classes normais
func (c *carro) inicializa(Nome, Marca, Modelo, Cor string) {
	c.Nome = Nome
	c.Marca = Marca
	c.Modelo = Modelo
	c.Cor = Cor
}
