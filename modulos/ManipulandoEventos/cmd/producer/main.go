package main

import "github.com/neberson/pos-go-expert-fullcycle/modulos/ManipulandoEventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, RabbitMQ!", "amq.direct")
}
