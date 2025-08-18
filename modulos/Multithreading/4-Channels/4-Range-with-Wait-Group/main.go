package main

import "sync"

// Tread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for value := range ch {
		println("Received:", value)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
