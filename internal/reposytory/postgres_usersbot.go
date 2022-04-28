package reposytory

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	adminTable = "admins"
)

type UsersBotRepository struct {
	db *sqlx.DB
}

func NewUsersBotRepository(db *sqlx.DB) *UsersBotRepository {
	return &UsersBotRepository{
		db: db,
	}
}

func (r *UsersBotRepository) GetAdmins(chatID int64, firstname string, lastname string) (string, error) {
	var id string
	qwery := fmt.Sprintf("select id from %s where first_name=($1) and last_name=($2)", adminTable)
	if err := r.db.Get(&id, qwery, firstname, lastname); err != nil {
		return "", err
	}
	return id, nil
}
