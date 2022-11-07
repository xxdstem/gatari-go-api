package user_db

import (
	"api/internal/entity"
	rep "api/internal/repository"
	"api/pkg/osuhelper"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) rep.UserRepository {
	return &repository{db: db}
}

func (r *repository) GetUsers(name string) ([]entity.User, error) {
	rows, err := r.db.Queryx("SELECT id, users.username, country, privileges, beta_key, email, username_aka FROM users LEFT JOIN users_stats USING (id) WHERE users.username LIKE ?", name)
	if err != nil {
		return nil, err
	}
	var result []entity.User
	for rows.Next() {
		u := entity.User{}
		if err := rows.StructScan(&u); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (r *repository) GetByID(id int) (*entity.User, error) {
	row := r.db.QueryRowx("SELECT id, users.username, ranked_maps, country, privileges, beta_key, email, username_aka, followers_count, play_style, favourite_mode FROM users LEFT JOIN users_stats USING (id) WHERE users.id = ?", id)
	result := entity.User{}
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetStatsByID(id int, mode int8) (*entity.UserStats, error) {
	q := "SELECT pp_%[1]s pp, level_%[1]s as 'level.level', total_hits_%[1]s total_hits, replays_watched_%[1]s replays_watched, playtime, max_combo_%[1]s max_combo, ranked_score_%[1]s as 'score.ranked', total_score_%[1]s AS 'score.total', avg_accuracy_%[1]s avg_accuracy, playcount_%[1]s play_count  FROM users_stats WHERE id = ?"
	row := r.db.QueryRowx(fmt.Sprintf(q, osuhelper.ModeToStr(mode)), id)
	result := entity.UserStats{}
	if err := row.StructScan(&result); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetRanks(id int, mode int8) (entity.Rankinkgs, error) {
	q := "SELECT COUNT(IF(rank = 'A', 1, null)) 'a_count', COUNT(IF(rank = 'S', 1, null)) s_count, COUNT(IF(rank = 'SH', 1, null)) sh_count, COUNT(IF(rank = 'X', 1, null)) x_count, COUNT(IF(rank = 'XH', 1, null)) xh_count FROM scores WHERE userid = ? AND play_mode = ? AND completed = 3"
	row := r.db.QueryRowx(q, id, mode)
	result := entity.Rankinkgs{}
	if err := row.StructScan(&result); err != nil {
		fmt.Println(err)
	}
	return result, nil
}
