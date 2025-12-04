package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) SignUp(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (telegram_username, first_name, last_name, grade, class_letter, password_hash, referral, by_referral, telegram_chat_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.TelegramUsername, user.FirstName, user.LastName, user.Grade, user.ClassLetter, user.PasswordHash, user.Referral, user.ByReferral, 0)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UsersPostgres) SignIn(telegramUsername string, passwordHash string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_username = $1 AND password_hash = $2", usersTable)
	if err := r.db.Get(&user, query, telegramUsername, passwordHash); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UsersPostgres) GetOneUser(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	if err := r.db.Get(&user, query, userId); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UsersPostgres) GetUserReferrals(userId int) ([]model.User, error) {
	user, err := r.GetOneUser(userId)
	if err != nil {
		return nil, err
	}
	var referralUsers []model.User
	query := fmt.Sprintf("SELECT first_name, last_name, telegram_username, grade, class_letter FROM %s WHERE by_referral = $1", usersTable)
	if err := r.db.Select(&referralUsers, query, user.Referral); err != nil {
		return nil, err
	}
	return referralUsers, nil
}
