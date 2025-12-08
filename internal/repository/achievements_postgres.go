package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type AchievementsPostgres struct {
	db *sqlx.DB
}

func NewAchievementsPostgres(db *sqlx.DB) *AchievementsPostgres {
	return &AchievementsPostgres{
		db: db,
	}
}

func (r *AchievementsPostgres) CreateAchievement(ach model.Achievement) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, node_id, coins, text, document_url) VALUES ($1, $2, $3, $4, $5) RETURNING id", achievementsTable)
	row := r.db.QueryRow(query, ach.UserId, ach.NodeId, ach.Coins, ach.Text, ach.DocumentUrl)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AchievementsPostgres) DeleteAchievement(achId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", achievementsTable)
	_, err := r.db.Exec(query, achId, userId)
	return err
}

func (r *AchievementsPostgres) GetUserAchievements(userId int) ([]model.Achievement, error) {
	var achs []model.Achievement
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", achievementsTable)
	if err := r.db.Select(&achs, query, userId); err != nil {
		return nil, err
	}
	return achs, nil
}

func (r *AchievementsPostgres) GetAchievementsByStatus(status string) ([]model.Achievement, error) {
	var achs []model.Achievement
	query := fmt.Sprintf("SELECT a.*, u.first_name AS user_first_name, u.last_name AS user_last_name, u.grade AS user_grade, u.class_letter AS user_class_letter FROM %s a JOIN %s u ON a.user_id = u.id WHERE a.status = $1", achievementsTable, usersTable)
	if err := r.db.Select(&achs, query, status); err != nil {
		return nil, err
	}
	return achs, nil
}

func (r *AchievementsPostgres) SetAchievementStatus(achId int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", achievementsTable)
	_, err := r.db.Exec(query, status, achId)
	return err
}

func (r *AchievementsPostgres) GetAchievementById(achId int) (model.Achievement, error) {
	var ach model.Achievement
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", achievementsTable)
	if err := r.db.Get(&ach, query, achId); err != nil {
		return model.Achievement{}, err
	}
	return ach, nil
}
