package main

import (
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/configs"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/infra/database"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/infra/webserver/handlers"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
