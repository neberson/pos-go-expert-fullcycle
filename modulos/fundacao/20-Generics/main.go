package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Neberson": 3000, "João": 2000, "Maria": 2500}
	m2 := map[string]float64{"Neberson": 3000.20, "João": 2000.10, "Maria": 2500.50}
	m3 := map[string]MyNumber{"Neberson": 3000, "João": 2000, "Maria": 2500}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
