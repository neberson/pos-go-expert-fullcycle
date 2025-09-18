package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, course CourseParams, category CategoryParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Price:       course.Price,
			CategoryID:  sql.NullString{String: category.ID, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, CourseID: %s, Course Name: %s, Course Description: %s, Course Price: %f", course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	/*courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Course 2",
		Description: sql.NullString{String: "Description Course 2", Valid: true},
		Price:       199.90,
	}

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Category 2",
		Description: sql.NullString{String: "Description Category 2", Valid: true},
	}

	CourseDB := NewCourseDB(dbConn)

	err = CourseDB.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	if err != nil {
		panic(err)
	}*/

}
