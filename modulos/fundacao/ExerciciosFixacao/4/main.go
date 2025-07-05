/*Exercício 4: Interfaces e Type Assertion
Enunciado:

Defina uma interface Speaker com um método Speak() string.

Crie dois tipos que implementem Speaker:

Dog cujo Speak() retorna "Woof!".

Cat cujo Speak() retorna "Meow!".

Escreva uma função MakeSpeak(s Speaker) que imprime o resultado de s.Speak().

No main, armazene numa slice de Speaker instâncias de Dog e Cat, e chame MakeSpeak para cada uma.

Bônus: dentro de MakeSpeak, detecte se s é *Dog usando type assertion e, nesse caso, imprima também "É um cachorro!"*/

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func MakeSpeak(s Speaker) {
	fmt.Printf("Speak: %s\n", s.Speak())

	if _, dog := s.(Dog); dog {
		fmt.Printf("É um cachorro\n")
	}
}

func main() {
	speakers := []Speaker{Dog{}, Cat{}}

	for _, v := range speakers {
		MakeSpeak(v)
	}
}
