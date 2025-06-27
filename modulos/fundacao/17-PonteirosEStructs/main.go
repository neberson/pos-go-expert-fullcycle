package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Neberson Andrade"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

type Conta struct {
	saldo int
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	neberson := Cliente{
		nome: "Neberson",
	}
	neberson.andou()
	fmt.Printf("O valor da struct com nome %v\n", neberson.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)

	conta.depositar(200)
	println(conta.saldo)
}
