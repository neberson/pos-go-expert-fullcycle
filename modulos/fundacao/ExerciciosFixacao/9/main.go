/*
Enunciado: Crie um novo tipo chamado Real baseado em float64 e outro tipo chamado Dolar também baseado em float64. Declare uma variável valorEmReais do tipo Real com o valor 50.0. Considere uma taxa de câmbio fixa de 1 Dólar = 4.95 Reais. Converta valorEmReais para Dolar e armazene o resultado em uma variável valorEmDolares do tipo Dolar. Imprima ambos os valores.
Objetivos de Aprendizado:
Criar tipos personalizados usando type.
Realizar conversão explícita entre tipos subjacentes iguais, mas tipos personalizados diferentes.
Compreender a importância de tipos personalizados para clareza e segurança.
*/
package main

import "fmt"

type Real float64
type Dolar float64

func main() {
	const taxaCambio = 4.95

	var valorEmReais Real = 50.0
	var valorEmDolares Dolar = Dolar(float64(valorEmReais) / taxaCambio)

	fmt.Printf("O valor em real é: %.2f\n", valorEmReais)
	fmt.Printf("O valor em dolar é: %.2f\n", valorEmDolares)
}
