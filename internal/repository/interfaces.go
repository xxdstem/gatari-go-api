package repository

import (
	"api/internal/entity"
)

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserStatsByID(id int, mode int8) (*entity.UserStats, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserRanks(id int, mode int8) (entity.Rankinkgs, error)
}

type UserRedisRepository interface {
	GetUserRank(user *entity.User, mode string) (entity.UserRank, error)
}

type BeatmapRepository interface {
	GetBeatmapByID(id int) (*entity.BeatmapSet, error)
}
