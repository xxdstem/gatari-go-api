package user_redis

import (
	"api/internal/entity"
	rep "api/internal/repository"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/redis.v5"
)

type repository struct {
	rd *redis.Client
}

func New(rd *redis.Client) rep.UserRedisRepository {
	return &repository{rd: rd}
}

func (r *repository) GetUserRank(user *entity.User, mode string) (entity.UserRank, error) {
	fmt.Println("HI PUTIN")
	rank := entity.UserRank{}
	global := r.rd.ZRevRank(fmt.Sprintf("ripple:leaderboard:%s", mode), strconv.Itoa(user.ID)).Val()
	country := r.rd.ZRevRank(fmt.Sprintf("ripple:leaderboard:%s:%s", mode, strings.ToLower(user.Country)), strconv.Itoa(user.ID)).Val()
	if global != 0 {
		rank.GlobalRank = int(global) + 1
		rank.CountryRank = int(country) + 1
	}

	return rank, nil
}
