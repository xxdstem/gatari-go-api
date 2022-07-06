package entity

type User struct {
	ID              int    `json:"id" db:"id"`
	Name            string `json:"username" db:"username"`
	NameAKA         string `json:"username_aka" db:"username_aka"`
	Privileges      int    `json:"-" db:"privileges"`
	BetaKey         string `json:"-" db:"beta_key"`
	Email           string `json:"-" db:"email"`
	RankedMapsCount int    `json:"ranked_maps" db:"ranked_maps"`
	Country         string `json:"country" db:"country"`
	FollowersCount  int    `json:"followers_count" db:"followers_count"`
	PlayTime        int    `json:"playtime" db:"playtime"`
	PlayStyle       int8   `json:"play_style" db:"play_style"`
	FavouriteMode   int8   `json:"favourite_mode" db:"favourite_mode"`
}

type UserStats struct {
	Name        string  `json:"username" db:"username"`
	PP          int     `json:"pp" db:"pp"`
	Level       int     `json:"level" db:"level"`
	TotalHits   int64   `json:"total_hits" db:"total_hits"`
	ReplayViews int     `json:"replays_watched" db:"replays_watched"`
	MaxCombo    int     `json:"max_combo" db:"max_combo"`
	RankedScore int64   `json:"ranked_score" db:"ranked_score"`
	TotalScore  int64   `json:"total_score" db:"total_score"`
	Accuracy    float64 `json:"avg_accuracy" db:"avg_accuracy"`
	PlayCount   int     `json:"play_count" db:"play_count"`
}
