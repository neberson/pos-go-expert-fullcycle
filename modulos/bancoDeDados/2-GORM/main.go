package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	//create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	product := Product{
		Name:       "Mouse",
		Price:      1000.0,
		CategoryID: 2,
	}

	db.Create(&product)
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
	// var products []Product

	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }

	// db.Create(&products)

	// //select one
	// var product Product
	// db.First(&product, 5)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// //select all
	// var products2 []Product
	// db.Find(&products2)
	// for _, product := range products2 {
	// 	fmt.Println(product)
	// }

	// var products3 []Product
	// db.Limit(2).Offset(3).Find(&products3)
	// for _, product := range products3 {
	// 	fmt.Println(product)
	// }

	// //where
	// var products4 []Product
	// db.Where("price > ?", 50).Find(&products4)
	// for _, product := range products4 {
	// 	fmt.Println(product)
	// }

	// //where
	// var products5 []Product
	// db.Where("name LIKE ?", "%book%").Find(&products5)
	// for _, product := range products5 {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// db.Delete(&p)
}
