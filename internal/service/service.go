package service

import "teleframBot/internal/reposytory"

type UsersBot interface {
	GetAdmins(chatID int64, firstname string, lastname string) (string, error)
}

type Service struct {
	UsersBot
}

func NewService(repo *reposytory.Reposytory) *Service {
	return &Service{
		UsersBot: NewUsersBotService(repo),
	}
}
