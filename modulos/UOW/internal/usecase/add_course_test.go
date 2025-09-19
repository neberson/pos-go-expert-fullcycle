package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/UOW/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS courses")
	dbt.Exec("DROP TABLE IF EXISTS categories")

	dbt.Exec("CREATE TABLE categories (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	dbt.Exec("CREATE TABLE courses (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), category_id INT, FOREIGN KEY (category_id) REFERENCES categories(id))")

	input := InputUseCase{
		CategoryName:     "Programação",
		CourseName:       "Go",
		CourseCategoryID: 2,
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
