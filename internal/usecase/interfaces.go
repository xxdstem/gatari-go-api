package usecase

import (
	"api/internal/entity"
)

type BeatmapsUseCase interface {
	UpdateBeatmapSet(id int) error
}
type UserUseCase interface {
	Update(id int) error
	GetStatsByID(id int, mode int8) *entity.UserStats
	GetById(id int) *entity.User
}

type LeaderboardUseCase interface {
}

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetStatsByID(id int, mode int8) (*entity.UserStats, error)
	GetByID(id int) (*entity.User, error)
	GetRanks(id int, mode int8) (entity.Rankinkgs, error)
}

type UserRedisRepository interface {
	GetUserRank(user *entity.User, mode string) (entity.UserRank, error)
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
