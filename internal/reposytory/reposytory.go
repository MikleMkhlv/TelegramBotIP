package reposytory

import "github.com/jmoiron/sqlx"

type UsersBot interface {
	GetAdmins(chatID int64, firstname string, lastname string) (string, error)
}

type Reposytory struct {
	UsersBot
}

func NewRepository(db *sqlx.DB) *Reposytory {
	return &Reposytory{
		UsersBot: NewUsersBotRepository(db),
	}
}
