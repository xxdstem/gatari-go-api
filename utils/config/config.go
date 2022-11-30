package cfg

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/xxdstem/gatari-go-api/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DSN           string
	Host          string
	APIKey        string
	RedisAddr     string
	RedisPassword string
	Workers       int
	DB            DB
}

type DB struct {
	Host     string
	DBname   string
	Port     string
	User     string
	Password string
}

var confFile = "./config.yml"

func NewConfig() (*Config, error) {
	log := logger.New()

	c := Config{
		DSN:       "root@/ripple",
		RedisAddr: "localhost:6379",
		Workers:   4,
		DB: DB{
			Host:     "127.0.0.1",
			DBname:   "ripple",
			Port:     "3306",
			User:     "",
			Password: "",
		},
	}

	f, err := os.Open(confFile)
	switch {
	case errors.Is(err, os.ErrNotExist):
		{
			log.Warn("No config.yml was found. Creating config.yml file...")

			f, err = os.Create(confFile)
			if errors.Is(err, os.ErrPermission) {
				log.Error("Can't create config.yml file, permission denied")
				return nil, err
			}
			if err != nil {
				log.Error("Can't create config.yml file")
				return nil, err
			}

			b, err := yaml.Marshal(&c)
			if err != nil {
				log.Error("Can't process config.yml")
				return nil, err
			}

			_, err = f.Write(b)
			if err != nil {
				log.Error("Can't write bytes into config.yml")
				return nil, err
			}
		}
	case err != nil:
		{
			log.Error(fmt.Sprintf("config.yml error: %v", err))
			return nil, err
		}
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Error("Can't read config.yml")
		return nil, err
	}

	yaml.Unmarshal(b, &c)

	// switch err := yaml.Load(&c, confFile); {
	// case err == conf.ErrNoFile:
	// 	log.Warn("No config.yml was found. Creating it")

	// 	err := conf.Export(&c, confFile)
	// 	if err != nil {
	// 		log.Error(fmt.Sprintf("Couldn't create config.conf: %v.", err))
	// 		return nil, err
	// 	}

	// 	log.Done("config.yml has been created!")
	// case err != nil:
	// 	log.Error(fmt.Sprintf("config.yml couldn't be loaded: %v.", err))
	// 	return nil, err
	// }
	return &c, nil
}
