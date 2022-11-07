package usecase

import (
	"api/internal/entity"
	"api/pkg/logging"
	"api/pkg/osuhelper"

	"zxq.co/ripple/ocl"
)

type _us struct {
	logger *logging.Logger
	rd     UserRedisRepository
	db     UserRepository
}

func NewUserUseCase(db UserRepository, rd UserRedisRepository, l *logging.Logger) UserUseCase {
	return &_us{
		logger: l,
		rd:     rd,
		db:     db,
	}
}

func (u *_us) GetById(id int) *entity.User {
	user, _ := u.db.GetByID(id)
	return user
}

func (u *_us) GetStatsByID(id int, mode int8) *entity.UserStats {
	user, err := u.db.GetByID(id)
	if err != nil {
		u.logger.Error(err)
		return nil
	}
	if mode == -1 {
		mode = user.FavouriteMode
	}
	userStats, _ := u.db.GetStatsByID(id, mode)
	if userStats.TotalHits > 0 && userStats.PlayCount > 0 {
		userStats.AvgHits = float64(userStats.TotalHits) / float64(userStats.PlayCount)
	}

	if rankings, err := u.db.GetRanks(id, mode); err == nil {
		userStats.Rankinkgs = rankings
	}
	modeStr := osuhelper.ModeToStr(mode)
	if userRank, err := u.rd.GetUserRank(user, modeStr); err == nil {
		userStats.Rank = userRank
	}
	userStats.Level.LevelProgress = osuhelper.LevelProgress(ocl.GetLevelPrecise(userStats.Score.TotalScore))

	return userStats
}

func (u *_us) Update(id int) error {
	return nil
}
