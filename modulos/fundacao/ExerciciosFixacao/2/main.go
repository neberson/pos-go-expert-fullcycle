/*Exercício 2: Mapas Básicos
Enunciado:

Crie um map[string]int para armazenar a idade de pessoas, com pelo menos três chaves: "Alice", "Bob" e "Carol".

Adicione mais uma pessoa ao mapa.

Verifique se existe a chave "Dave"; se não existir, insira "Dave" com valor 30.

Finalmente, imprima todas as chaves e valores do mapa.

Objetivos de aprendizado:

Declarar e inicializar map.

Inserir, verificar existência e iterar em maps.*/

package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 20, "Bob": 30, "Carol": 25}

	ages["Eve"] = 35

	if _, exists := ages["Dave"]; !exists {
		ages["Dave"] = 30
	}

	for name, age := range ages {
		fmt.Printf("%s tem %d anos\n", name, age)
	}
}
