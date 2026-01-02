package model

type Achievement struct {
	Id              int    `json:"id" db:"id"`
	UserId          int    `json:"user_id" db:"user_id"`
	NodeId          int    `json:"node_id" db:"node_id"`
	Coins           int    `json:"coins" db:"coins"`
	Text            string `json:"text" db:"text"`
	DocumentUrl     string `json:"document_url" db:"document_url"`
	Status          string `json:"status" db:"status"`
	UserFirstName   string `json:"user_first_name" db:"user_first_name"`
	UserLastName    string `json:"user_last_name" db:"user_last_name"`
	UserGrade       int    `json:"user_grade" db:"user_grade"`
	UserClassLetter string `json:"user_class_letter" db:"user_class_letter"`
	RejectComment   string `json:"reject_comment" db:"reject_comment"`
}
