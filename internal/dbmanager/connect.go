package dbmanager

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type db struct {
	Db       *gorm.DB
	login    string
	password string
}

func Connect(login string, password string) (*db, error) {
	dsn := login + ":" + password + "@(163.172.98.147)/ripple?charset=utf8mb4"
	ddb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &db{
		Db: ddb,
	}, nil
}
