package main

import "fmt"

const a = "Hello, World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "neberson"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("\nO valor de E é %v", e)
	fmt.Printf("\nO tipo de f é %T", f)
}

func x() {
}
