package main

import (
	"fmt"
	"strings"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func main() {
	cliente := Cliente{
		Nome:  "Vagner",
		Email: "v@v.com",
		CPF:   1212121212,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"Bia", "bia@bia.com", 44557798}
	fmt.Println(cliente2)
	cliente2.Nome = Upper(cliente2.Nome)

	fmt.Println(cliente2.Nome)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Upperx(s *string) {
	var z string
	z = s
	fmt.Printf(z)

}
