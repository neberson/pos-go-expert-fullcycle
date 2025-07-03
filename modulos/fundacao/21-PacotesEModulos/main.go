package main

import (
	"fmt"

	"PacotesEModulos/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())
	fmt.Println(carro)
	fmt.Println("Resultado: ", soma)
	fmt.Println(matematica.A)
}
