package database

import "github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
