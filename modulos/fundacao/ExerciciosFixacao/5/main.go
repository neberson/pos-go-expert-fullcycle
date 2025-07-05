/* Exercício 5: Closures e Funções Variádicas
Enunciado:

Implemente uma função SumGenerator() que retorna duas funções (closures):

Add(nums ...int) que adiciona os valores de nums ao acumulador interno.

Total() int que retorna o valor acumulado até o momento.

No main, crie um par dessas funções, chame Add(1,2,3), depois Add(4,5), e finalmente imprima Total().

Objetivos de aprendizado:

Criar closures que compartilham estado interno.

Usar funções variádicas (...int). */

package main

import "fmt"

func SumGenerator() (func(nums ...int), func() int) {
	sum := 0
	add := func(nums ...int) {
		for _, num := range nums {
			sum += num
		}
	}

	total := func() int {
		return sum
	}

	return add, total
}

func main() {
	add, total := SumGenerator()

	add(1, 2, 3)
	add(4, 5)

	fmt.Printf("O total dos valores é: %d", total())
}
