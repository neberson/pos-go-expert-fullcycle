package main

// Thread 1
func main() {
	canal := make(chan string) // Vazio

	// Thread 2
	go func() {
		canal <- "Olá, canal!" // Está cheio
	}()

	// Thread 1
	msg := <-canal // Canal esvazia
	println(msg)
}
