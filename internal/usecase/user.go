package usecase

import (
	"api/internal/entity"
	"api/pkg/logging"
)

func NewUserUseCase(db UserRepository, l *logging.Logger) UserUseCase {
	return &_userUseCase{
		logger: l,
		db:     db,
	}
}

func (u *_userUseCase) GetUserById(id int) *entity.User {
	user, _ := u.db.GetUserByID(id)
	return user
}

func (u *_userUseCase) UpdateUser(id int) error {
	return nil
}
