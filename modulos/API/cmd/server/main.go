package main

import "github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
