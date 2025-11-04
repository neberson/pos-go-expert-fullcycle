package main

func add(a, b int) int {
	sum := a + b // sum é alocado na stack
	return sum   // liberada quando a função termina
}

type User struct {
	Name string
}

func main() {
	println(add(1, 2))

	var p *int
	{
		x := 42
		p = &x // x escapa para o heap, pois seu endereço é retornado
	}
}
