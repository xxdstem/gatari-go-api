package usecase

import (
	"api/internal/entity"
	"api/pkg/logging"
	"api/pkg/osuhelper"

	"zxq.co/ripple/ocl"
)

func NewUserUseCase(db UserRepository, rd UserRedisRepository, l *logging.Logger) UserUseCase {
	return &_userUseCase{
		logger: l,
		rd:     rd,
		db:     db,
	}
}

func (u *_userUseCase) GetUserById(id int) *entity.User {
	user, _ := u.db.GetUserByID(id)
	return user
}

func (u *_userUseCase) GetUserStatsByID(id int, mode int8) *entity.UserStats {
	user, err := u.db.GetUserByID(id)
	if err != nil {
		u.logger.Error(err)
		return nil
	}
	if mode == -1 {
		mode = user.FavouriteMode
	}
	userStats, _ := u.db.GetUserStatsByID(id, mode)
	userStats.AvgHits = float64(userStats.TotalHits / int64(userStats.PlayCount))
	if rankings, err := u.db.GetUserRanks(id, mode); err == nil {
		userStats.Rankinkgs = rankings
	}
	modeStr := osuhelper.ModeToStr(mode)
	if userRank, err := u.rd.GetUserRank(user, modeStr); err == nil {
		userStats.Rank = userRank
	}
	userStats.Level.LevelProgress = osuhelper.LevelProgress(ocl.GetLevelPrecise(userStats.Score.TotalScore))

	return userStats
}

func (u *_userUseCase) UpdateUser(id int) error {
	return nil
}
