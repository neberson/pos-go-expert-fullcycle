/*Exercício 3: Structs, Métodos e Ponteiros
Enunciado:

Defina um struct chamado Rectangle com os campos Width e Height (float64).

Implemente dois métodos:

Area() float64 recebe valor (receiver não ponteiro) e retorna a área.

Scale(factor float64) recebe ponteiro e multiplica os campos Width e Height pelo fator.

No main, crie um Rectangle, escale-o por 1.5 e imprima antes e depois a área.

Objetivos de aprendizado:

Composição de structs e definição de métodos.

Diferença entre receiver por valor e por ponteiro.*/

package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Height *= factor
	r.Width *= factor
}

func main() {
	rectangle := Rectangle{Width: 2, Height: 3}

	fmt.Printf("Área original: %f\n", rectangle.Area())

	rectangle.Scale(1.5)

	fmt.Printf("Área escalada:: %f\n", rectangle.Area())
}
