package app

import (
	"api/config"
	httpHandler "api/internal/controller/http"
	redisHandler "api/internal/controller/redis"
	"api/internal/repository/beatmap_db"
	"api/internal/repository/user_db"
	"api/internal/repository/user_redis"
	"api/pkg/logging"
	"net"
	"net/http"
	"time"

	"api/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/redis.v5"
)

type Router struct {
	j *httprouter.Router
	l *logging.Logger
}

func (c Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	header := w.Header()
	if r.Method == http.MethodOptions {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))

		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	}
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")
	c.j.ServeHTTP(w, r)
	c.l.Debugf("> Request end - time took: %s", time.Since(begin).String())
}

func Run(conf *config.Config, logger *logging.Logger) {
	logger.Info("Started!")
	db, err := sqlx.Open("mysql", conf.DSN)

	if err != nil || db.Ping() != nil {
		logger.Fatalf("couldn't start MySQL connection: %v.", err)
		return
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)
	r := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
	})

	usersUseCase := usecase.NewUserUseCase(
		user_db.New(db),
		user_redis.New(r),
		logger,
	)
	beatmapsUseCase := usecase.NewBeatmapsUseCase(
		beatmap_db.New(db),
		logger,
	)

	// TODO: implement repos
	leaderboardUseCase := usecase.NewLeaderboardUseCase(
		leaderboard_db.New(db),
		leaderboard_redis.New(r),
		logger,
	)

	router := httprouter.New()
	httpHandler.NewUsersRoute(router, logger, usersUseCase, beatmapsUseCase)
	httpHandler.NewLeaderboardRoute(router, logger, leaderboardUseCase)

	enableCors(router)
	jsonRouter := &Router{router, logger}

	redisHandler.NewRouter(r, logger, usersUseCase, beatmapsUseCase)

	logger.Info("starting web!")
	startWeb(jsonRouter)

}

func enableCors(router *httprouter.Router) {
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
}

func startWeb(router *Router) {
	l, err := net.Listen("tcp", ":8008")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: router,
	}
	panic(server.Serve(l))
}
