/*Enunciado: Declare múltiplas variáveis em uma única linha para armazenar o nome (string), idade (int) e altura (float64) de uma pessoa. Atribua valores a essas variáveis. Em seguida, imprima todas as informações em uma única linha formatada, por exemplo: "Nome: João, Idade: 30, Altura: 1.75m".
Objetivos de Aprendizado:
Declarar múltiplas variáveis de diferentes tipos.
Atribuir valores a múltiplas variáveis.
Usar fmt.Printf com uma combinação de verbos de formatação (%s, %d, %.2f).*/

package main

import "fmt"

func main() {
	nome, idade, altura := "Neberson", 35, 1.86
	fmt.Printf("Nome: %s, Idade: %d, Altura: %.2fm", nome, idade, altura)
}
