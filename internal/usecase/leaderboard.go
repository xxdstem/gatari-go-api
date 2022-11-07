package usecase

import (
	"api/pkg/logging"
)

type _lb struct {
	logger *logging.Logger
	rd     UserRedisRepository
	db     UserRepository
}

func NewLeaderboardUseCase(db UserRepository, rd UserRedisRepository, l *logging.Logger) LeaderboardUseCase {
	return &_lb{
		logger: l,
		rd:     rd,
		db:     db,
	}
}

func (l *_lb) Global() {
	// TODO: Get Global Leaderboard
}
