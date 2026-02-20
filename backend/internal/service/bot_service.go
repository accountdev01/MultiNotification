package service

import (
	"github.com/accountdev01/MultiNotification.git/internal/model"
	"github.com/accountdev01/MultiNotification.git/internal/repository"
)

type BotService struct {
	repo repository.BotRepository
}

func NewBotService(repo repository.BotRepository) *BotService {
	return &BotService{repo: repo}
}

func (s *BotService) CreateBot(bot *model.Bot) error {
	return s.repo.CreateBot(bot)
}

func (s *BotService) GetAllBots() ([]model.Bot, error) {
	return s.repo.GetAllBots()
}

func (s *BotService) GetBotByID(id string) (*model.Bot, error) {
	return s.repo.GetBotByID(id)
}
