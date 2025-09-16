package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = executeCreateCategory(ctx, queries)
	if err != nil {
		panic(err)
	}

	err = executeUpdateCategory(ctx, queries)
	if err != nil {
		panic(err)
	}

	err = executeDeleteCategory(ctx, queries)
	if err != nil {
		panic(err)
	}
}

func executeCreateCategory(ctx context.Context, queries *db.Queries) error {

	log.Println("-------- Creating categories --------")

	err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Category 3",
		Description: sql.NullString{String: "Description 3", Valid: true},
	})

	if err != nil {
		return err
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		return err
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	return nil
}

func executeUpdateCategory(ctx context.Context, queries *db.Queries) error {

	log.Println("-------- Updating categories --------")

	err := queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "e20a177a-cd36-4923-bf8f-df5ff7bea3f7",
		Name:        "Category 1 - updated",
		Description: sql.NullString{String: "Description 1 - updated", Valid: true},
	})

	if err != nil {
		return err
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		return err
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	return nil
}

func executeDeleteCategory(ctx context.Context, queries *db.Queries) error {

	log.Println("-------- Deleting categories --------")

	err := queries.DeleteCategory(ctx, "1c4e5ac5-a32c-4bef-98df-21d741f4b60c")
	if err != nil {
		return err
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		return err
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	return nil
}
