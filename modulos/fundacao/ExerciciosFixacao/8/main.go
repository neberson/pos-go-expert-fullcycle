/*
Enunciado: Crie um programa Go que declare duas variáveis inteiras, num1 e num2, e atribua a elas os valores 15 e 7, respectivamente. Em seguida, declare uma terceira variável soma e armazene nela o resultado da adição de num1 e num2. Finalmente, imprima o resultado no formato "A soma de [num1] e [num2] é [soma]".
Objetivos de Aprendizado:
Declarar e atribuir variáveis usando a declaração curta (:=).
Realizar operações aritméticas básicas.
Usar fmt.Printf com múltiplos verbos de formatação.
*/
package main

import "fmt"

func main() {
	num1 := 15
	num2 := 7
	soma := num1 + num2

	fmt.Printf("A soma de %d e %d é %d", num1, num2, soma)
}
