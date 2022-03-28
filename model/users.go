package model

import "time"

type Users struct {
	Id        string
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsDeleted bool      `db:"is_deleted"`
}
