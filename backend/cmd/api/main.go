package main

import (
	"github.com/accountdev01/MultiNotification.git/internal/database"
	"github.com/accountdev01/MultiNotification.git/internal/handler"
	"github.com/accountdev01/MultiNotification.git/internal/model"
	"github.com/accountdev01/MultiNotification.git/internal/repository"
	"github.com/accountdev01/MultiNotification.git/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	db.AutoMigrate(&model.Bot{})

	botRepo := repository.NewBotRepository(db)
	botSvc := service.NewBotService(botRepo)
	botHdl := handler.NewBotHandler(botSvc)

	// 4. Setup Router
	r := gin.Default()
	botHdl.RegisterRoutes(r)

	r.Run(":8080")
}
