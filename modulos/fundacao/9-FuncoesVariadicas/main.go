package main

import "fmt"

func main() {
	fmt.Println(sum(123, 456, 1846, 84, 21, 584, 546, 45, 4756, 4654, 4, 83))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
