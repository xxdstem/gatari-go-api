package utils

import (
	"fmt"

	"github.com/thehowl/conf"
	"github.com/xxdstem/gatari-go-api/logger"
)

type Config struct {
	DSN           string
	Host          string
	APIKey        string
	RedisAddr     string
	RedisPassword string
	Workers       int
}

var confFile = "./config/config.conf"

func NewConfig() (*Config, error) {
	log := logger.New()

	c := Config{
		DSN:       "root@/ripple",
		RedisAddr: "localhost:6379",
		Workers:   4,
	}

	switch err := conf.Load(&c, confFile); {
	case err == conf.ErrNoFile:
		log.Warn("No config.conf was found. Creating it")
		err := conf.Export(&c, confFile)
		if err != nil {
			log.Error(fmt.Sprintf("Couldn't create config.conf: %v.", err))
			return nil, err
		} else {
			log.Done("config.conf has been created!")
		}
	case err != nil:
		log.Error(fmt.Sprintf("config.conf couldn't be loaded: %v.", err))
		return nil, err
	}
	return &c, nil
}
