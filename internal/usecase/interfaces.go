package usecase

import (
	"api/internal/entity"
)

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}
