package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type message struct {
	MsgId int64
	Msg   string
}

func main() {
	c1 := make(chan message)
	c2 := make(chan message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabitMQ: ID: %d-%s\n", msg.MsgId, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Received from Kafka: ID: %d-%s\n", msg.MsgId, msg.Msg)
		case <-time.After(3 * time.Second):
			println("timeout")
		}
	}
}
