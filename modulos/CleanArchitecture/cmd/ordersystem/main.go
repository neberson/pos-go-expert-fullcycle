package main

import (
	"database/sql"
	"fmt"

	config "github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/configs"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/event/handler"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/pkg/events"
	"github.com/streadway/amqp"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()
	defer rabbitMQChannel.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
