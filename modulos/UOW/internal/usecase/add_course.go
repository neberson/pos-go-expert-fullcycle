package usecase

import (
	"context"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/UOW/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/UOW/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(
	courseRepo repository.CourseRepositoryInterface, categoryRepo repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:   courseRepo,
		CategoryRepository: categoryRepo,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
