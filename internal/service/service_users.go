package service

import (
	"teleframBot/internal/reposytory"
)

type UsersBotService struct {
	repo reposytory.UsersBot
}

func NewUsersBotService(repo reposytory.UsersBot) *UsersBotService {
	return &UsersBotService{
		repo: repo,
	}
}

func (s *UsersBotService) GetAdmins(chatID int64, firstname string, lastname string) (string, error) {
	return s.repo.GetAdmins(chatID, firstname, lastname)
}
