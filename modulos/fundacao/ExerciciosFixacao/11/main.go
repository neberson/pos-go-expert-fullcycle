/*
Enunciado: Crie dois novos tipos: Centimetros e Metros, ambos baseados em float64.
Declare uma variável distanciaCm do tipo Centimetros e atribua a ela o valor 250.0.
Crie uma função chamada ConverterParaMetros que receba um valor do tipo Centimetros e retorne o valor equivalente do tipo Metros.
Dentro da função main, chame ConverterParaMetros com distanciaCm e imprima o resultado.
Desafio Extra: Tente criar uma função ConverterParaCentimetros que faça o inverso.
Objetivos de Aprendizado:
Aprofundar o uso de tipos personalizados.
Criar e usar funções que recebem e retornam tipos personalizados.
Reforçar a necessidade de conversão explícita entre tipos personalizados, mesmo com o mesmo tipo subjacente.
*/
package main

import "fmt"

type Centimentros float64
type Metros float64

func ConverterParaMetros(cm Centimentros) Metros {
	return Metros(float64(cm) / 100.0)
}

func ConverterParaCentimetros(m Metros) Centimentros {
	return Centimentros(float64(m) * 100.0)
}

func main() {
	distanciaCm := Centimentros(250.0)
	distanciaM := ConverterParaMetros(distanciaCm)
	fmt.Printf("%.2f Centímetros equivalem a %.2f Metros.\n", distanciaCm, distanciaM)

	distanciaM2 := Metros(2.5)
	distanciaCm2 := ConverterParaCentimetros(distanciaM2)
	fmt.Printf("%.2f Metros equivalem a %.2f Centímetros.\n", distanciaM2, distanciaCm2)
}
