/* Exercício 1: Iterando Arrays e Slices
Enunciado:

Declare um array de inteiros com 5 elementos e inicialize-o com os valores 10, 20, 30, 40, 50.

Converta esse array em um slice que contenha apenas os três primeiros elementos.

Imprima no console cada elemento do slice, precedido do seu índice.*/

package main

func main() {
	var arr = [5]int{10, 20, 30, 40, 50}

	subslice := arr[0:3]

	for i, v := range subslice {
		println(i, v)
	}
}
