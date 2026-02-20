package repository

import (
	"github.com/accountdev01/MultiNotification.git/internal/model"
	"gorm.io/gorm"
)

type BotRepository interface {
	CreateBot(bot *model.Bot) error
	GetBotByID(id string) (*model.Bot, error)
	GetAllBots() ([]model.Bot, error)
}

type botRepo struct {
	db *gorm.DB
}

func NewBotRepository(db *gorm.DB) BotRepository {
	return &botRepo{db: db}
}

func (r *botRepo) CreateBot(bot *model.Bot) error {
	return r.db.Create(bot).Error
}

func (r *botRepo) GetAllBots() ([]model.Bot, error) {
	var bots []model.Bot
	if err := r.db.Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *botRepo) GetBotByID(id string) (*model.Bot, error) {
	var bot model.Bot
	if err := r.db.Where("id = ?", id).First(&bot).Error; err != nil {
		return nil, err
	}
	return &bot, nil
}
