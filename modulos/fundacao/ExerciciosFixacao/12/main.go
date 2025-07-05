/*
Enunciado: Crie um array de strings com tamanho fixo de 4 para representar uma lista de compras. Atribua os seguintes itens: "Pão", "Leite", "Ovos", "Café". Em seguida, imprima o array completo e o item na segunda posição (índice 1).
Objetivos de Aprendizado:
Declarar e inicializar arrays.
Acessar elementos de um array por índice.
Compreender a natureza de tamanho fixo dos arrays.
*/
package main

import "fmt"

func main() {
	var listaCompras [4]string = [4]string{"Pão", "Leite", "Ovos", "Café"}

	fmt.Println("Lista de Compras:", listaCompras)
	fmt.Println("Item na segunda posição (índice 1):", listaCompras[1])
}
