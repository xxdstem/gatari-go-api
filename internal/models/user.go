package models

type User struct {
	ID                  int
	Username            string `gorm:"type:varchar(32);collate:utf8_general_ci"`
	UsernameSafe        string `gorm:"type:varchar(32);collate:utf8mb4_unicode_ci"`
	Password            string `gorm:"type:varchar(2048);collate:latin1_swedish_ci"`
	PasswordMd5         string `gorm:"type:varchar(64);collate:latin1_swedish_ci"`
	Vk                  int
	Salt                string `gorm:"type:varchar(32);collate:utf8_bin"`
	PasswordVersion     int
	Privileges          int
	Allowed             int
	Aqn                 int
	Flags               int
	DonorExpire         int
	AchievementsVersion int    `gorm:"type:tinyint(1)"`
	SilenceReason       string `gorm:"type:varchar(2048);collate:cp1251_bin"`
	SilenceEnd          int
	BanDatetime         int
	Latest_Activity     int
	BetaKey             string `gorm:"type:varchar(32);collate:utf8_general_ci"`
	Notes               string `gorm:"collate:utf8mb4_unicode_ci"`
	RegisterDatetime    int
	Email               string `gorm:"type:varchar(128);collate:latin1_swedish_ci"`
	Land                string `gorm:"type:varchar(4);collate:latin1_swedish_ci"`
}
