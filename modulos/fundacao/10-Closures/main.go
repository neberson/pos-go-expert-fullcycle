package main

import "fmt"

func main() {

	total := func() int {
		return sum(123, 456, 1846, 84, 21, 584, 546, 45, 4756, 4654, 4, 83) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
