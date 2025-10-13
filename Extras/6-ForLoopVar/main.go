package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Println(i)
	}

	done := make(chan bool)
	values := []string{"A", "B", "C"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
