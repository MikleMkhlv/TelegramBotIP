package reposytory

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgresql://postgres:%s@%s/%s?%s", cfg.Password, cfg.Host, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
