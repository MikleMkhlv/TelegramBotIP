package reposytory

import "github.com/jmoiron/sqlx"

type UsersBot interface {
}

type Reposytory struct {
	UsersBot
}

func NewReposytory(db *sqlx.DB) *Reposytory {
	return &Reposytory{}
}
