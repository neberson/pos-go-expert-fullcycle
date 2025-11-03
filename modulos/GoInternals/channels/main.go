package main

import (
	"fmt"
	"time"
)

func main() {
	simpleChannel()
	bufferedChannel()
	bufferedChannelBloq()
	notBufferedChannelTimeout()
	notBufferedChannelWorkers()
}

func simpleChannel() {
	ch := make(chan int)

	go func() {
		ch <- 42 // envia valor para o channel
	}()

	value := <-ch      // recebe valor do channel
	fmt.Println(value) // imprime 42
}

func bufferedChannel() {
	ch := make(chan int, 2) // canal com buffer de tamanho 2

	ch <- 42
	ch <- 43

	fmt.Println(<-ch) // imprime 42
	fmt.Println(<-ch) // imprime 43
}

func bufferedChannelBloq() {
	ch := make(chan int, 2) // canal com buffer de tamanho 2

	go func() {
		ch <- 42
		fmt.Println("Sent 1")
		ch <- 43
		fmt.Println("Sent 2")
		ch <- 44 // bloqueia aqui, pois o buffer está cheio
		fmt.Println("Sent 3")
	}()

	time.Sleep(time.Second)

	go func() {
		val := <-ch
		fmt.Println("Received", val)
	}()

	time.Sleep(2 * time.Second)
}

func notBufferedChannelTimeout() {
	ch := make(chan int) // canal sem buffer
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42 // bloqueia aqui até que alguém receba o valor
	}()

	select {
	case val := <-ch:
		fmt.Println("Received", val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: no value received")
	}
}

func notBufferedChannelWorkers() {
	done := make(chan bool)
	numWorkers := 3

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Worker %d done\n", id)
			done <- true
		}(i)
	}

	for i := 0; i < numWorkers; i++ {
		<-done // espera cada worker terminar
	}

	fmt.Println("All workers done")
}
