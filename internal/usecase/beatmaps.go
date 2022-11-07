package usecase

import "api/pkg/logging"

type _b struct {
	logger *logging.Logger
	db     BeatmapRepository
}

func NewBeatmapsUseCase(db BeatmapRepository, l *logging.Logger) BeatmapsUseCase {
	return &_b{
		logger: l,
		db:     db,
	}
}

func (u *_b) GetBeatmapByID(id int) error {
	_, err := u.db.GetBeatmapByID(id)
	if err != nil {
		return err
	}
	return nil
}
