package model

type Node struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Points   int    `json:"points"`
	ParentID int64  `json:"parent_id,omitempty"`
}
