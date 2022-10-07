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
	PlayStyle       int8   `json:"play_style" db:"play_style"`
	FavouriteMode   int8   `json:"favourite_mode" db:"favourite_mode"`
}

type UserRank struct {
	GlobalRank  int `json:"global"`
	CountryRank int `json:"country"`
}

type Rankinkgs struct {
	X  int `json:"x" db:"x_count"`
	XH int `json:"xh" db:"xh_count"`
	S  int `json:"s" db:"s_count"`
	SH int `json:"sh" db:"sh_count"`
	A  int `json:"a" db:"a_count"`
}
type UserStats struct {
	Rank     UserRank `json:"rank"`
	PlayTime int      `json:"playtime" db:"playtime"`
	Level    struct {
		Level         int     `json:"level" db:"level"`
		LevelProgress float64 `json:"progress"`
	} `json:"level"`

	Rankinkgs Rankinkgs `json:"rankinkgs"`

	Score struct {
		RankedScore int64 `json:"ranked" db:"ranked"`
		TotalScore  int64 `json:"total" db:"total"`
	} `json:"score" db:"score"`

	AvgHits     float64 `json:"avg_hits"`
	PP          int     `json:"pp" db:"pp"`
	TotalHits   int64   `json:"total_hits" db:"total_hits"`
	ReplayViews int     `json:"replays_watched" db:"replays_watched"`
	MaxCombo    int     `json:"max_combo" db:"max_combo"`
	Accuracy    float64 `json:"avg_accuracy" db:"avg_accuracy"`
	PlayCount   int     `json:"play_count" db:"play_count"`
}
