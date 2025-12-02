package model

type User struct {
	Id int `json:"id" db:"id"`
	TelegramUsername string `json:"telegram_username" db:"telegram_username"`
	TelegramChatId int `json:"telegram_chat_id" db:"telegram_chat_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Grade int `json:"grade" db:"grade"`
	ClassLetter string `json:"class_letter" db:"class_letter"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Referral string `json:"referral" db:"referral"`
	ByReferral string `json:"by_referral" db:"by_referral"`
}

