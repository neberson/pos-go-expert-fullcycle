/*Enunciado: Crie um programa Go que declare uma variável nome do tipo string e atribua a ela o seu nome. Em seguida, use a função fmt.Printf para imprimir uma saudação personalizada no console, como "Olá, [Seu Nome]! Bem-vindo(a) ao Go!".
Objetivos de Aprendizado:
Entender a estrutura básica de um programa Go (package main, func main()).
Declarar e atribuir valores a variáveis usando var.
Importar e usar o pacote fmt para saída formatada.*/

package main

import "fmt"

func main() {
	var nome string = "Neberson"
	fmt.Printf("Olá, %s! Bem-vindo(a) ao Go!", nome)
}
