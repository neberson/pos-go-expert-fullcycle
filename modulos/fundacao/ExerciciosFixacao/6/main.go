/*Exercício 6: Generics com Tipos Customizados
Enunciado:

Implemente uma função genérica Filter[T any](slice []T, keep func(T) bool) []T que retorna um novo slice contendo apenas os elementos que satisfazem keep.

No main, use Filter para:

Filtrar um slice de int, mantendo apenas números pares.

Filtrar um slice de string, mantendo apenas palavras com mais de 3 caracteres.

Imprima os resultados.

Objetivos de aprendizado:

Criar funções genéricas usando [T any].

Passar callbacks (funções) para genéricos.*/

package main

import "fmt"

func Filter[T any](sliceGenerico []T, keep func(T) bool) []T {
	t := []T{}
	for _, valor := range sliceGenerico {
		if keep(valor) {
			t = append(t, valor)
		}
	}
	return t
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	evens := Filter(nums, func(valor int) bool {
		return valor%2 == 0
	})
	fmt.Println("Pares:", evens)

	nomes := []string{"Neberson", "João", "Maria", "de", "di", "did"}
	nomesFiltrados := Filter(nomes, func(nome string) bool {
		return len(nome) > 3
	})
	fmt.Println("Nomes >3 letras:", nomesFiltrados)
}
