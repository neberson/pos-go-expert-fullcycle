package main

import "fmt"

func main() {
	salarios := map[string]int{"Neberson": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Neberson"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	sal := make(map[string]int)
	fmt.Println(sal)

	sal1 := map[string]int{}

	fmt.Println(sal1)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
}
