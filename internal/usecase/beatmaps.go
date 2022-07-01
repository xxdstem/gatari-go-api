package usecase

import "api/pkg/logging"

func NewBeatmapsUseCase(db BeatmapRepository, l *logging.Logger) BeatmapsUseCase {
	return &_beatmapUseCase{
		logger: l,
		db:     db,
	}
}

func (u *_beatmapUseCase) UpdateBeatmapSet(id int) error {
	u.logger.Info("requested updating beatmapset ", id)
	_, err := u.db.GetBeatmapByID(id)
	if err != nil {
		return err
	}
	return nil
}
