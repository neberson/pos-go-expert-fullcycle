package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	for _, v := range numeros {
		println(v)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Hello, world")
	}

}
