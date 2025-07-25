package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.", c.Nome)
}

func main() {
	neberson := Cliente{
		Nome:  "Neberson",
		Idade: 38,
		Ativo: true,
	}
	neberson.Cidade = "Goiania"
	fmt.Println(neberson)

	neberson.Ativo = false
	fmt.Println(neberson.Ativo)

	neberson.Desativar()
}
