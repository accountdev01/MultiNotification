package handler

import (
	"github.com/accountdev01/MultiNotification.git/internal/model"
	"github.com/accountdev01/MultiNotification.git/internal/service"
	"github.com/gin-gonic/gin"
)

type BotHandler struct {
	svc *service.BotService
}

func NewBotHandler(svc *service.BotService) *BotHandler {
	return &BotHandler{svc: svc}
}

func (h *BotHandler) RegisterRoutes(r *gin.Engine) {
	botGroup := r.Group("/api/bots")
	{
		botGroup.POST("/add", h.CreateBot)
		botGroup.GET("/:id", h.GetBotByID)
		botGroup.GET("/", h.GetAllBots)
	}
}

func (h *BotHandler) CreateBot(c *gin.Context) {
	var bot model.Bot
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.svc.CreateBot(&bot); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create bot"})
		return
	}

	c.JSON(201, bot)
}

func (h *BotHandler) GetAllBots(c *gin.Context) {
	bots, err := h.svc.GetAllBots()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve bots"})
		return
	}
	c.JSON(200, bots)
}

func (h *BotHandler) GetBotByID(c *gin.Context) {
	id := c.Param("id")
	bot, err := h.svc.GetBotByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Bot not found"})
		return
	}
	c.JSON(200, bot)
}
