package model

type NevermorePlugin struct {
	ID      string `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Type    string `db:"type" json:"type"`
	OwnerID string `db:"owner"`
}
