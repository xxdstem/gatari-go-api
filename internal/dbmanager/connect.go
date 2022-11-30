package dbmanager

import (
	"strings"

	cfg "github.com/xxdstem/gatari-go-api/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type db struct {
	Db *gorm.DB
}

func Connect(c *cfg.DB) (*db, error) {
	var dsn strings.Builder

	dsn.WriteString(c.User)
	dsn.WriteString(":")
	dsn.WriteString(c.Password)
	dsn.WriteString("@(")
	dsn.WriteString(c.Host)
	dsn.WriteString(":")
	dsn.WriteString(c.Port)
	dsn.WriteString(")/")
	dsn.WriteString(c.DBname)
	dsn.WriteString("?charset=utf8mb4")

	ddb, err := gorm.Open(mysql.Open(dsn.String()))
	if err != nil {
		return nil, err
	}

	return &db{
		Db: ddb,
	}, nil
}
