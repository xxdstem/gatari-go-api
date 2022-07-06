package usecase

import (
	"api/internal/entity"
	"api/pkg/logging"
)

type _userUseCase struct {
	logger *logging.Logger
	db     UserRepository
}
type _beatmapUseCase struct {
	logger *logging.Logger
	db     BeatmapRepository
}

type BeatmapsUseCase interface {
	UpdateBeatmapSet(id int) error
}
type UserUseCase interface {
	UpdateUser(id int) error
	GetUserStatsByID(id int, mode int8) *entity.UserStats
	GetUserById(id int) *entity.User
}

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserStatsByID(id int, mode int8) (*entity.UserStats, error)
	GetUserByID(id int) (*entity.User, error)
}

type UserMeiliRepository interface {
	UpdateUser(*entity.User) error
}

type BeatmapRepository interface {
	GetBeatmapByID(id int) (*entity.BeatmapSet, error)
}

type BeatmapMeiliRepository interface {
	UpdateBeatmap(*entity.BeatmapSet) error
}
