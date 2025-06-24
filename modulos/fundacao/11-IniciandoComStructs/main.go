package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	neberson := Cliente{
		Nome:  "Neberson",
		Idade: 38,
		Ativo: true,
	}
	fmt.Println(neberson)

	neberson.Ativo = false
	fmt.Println(neberson.Ativo)
}
