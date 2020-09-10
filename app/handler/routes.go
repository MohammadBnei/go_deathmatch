package handler

import (
	"esgi_go/game"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(e *gin.Engine, channel chan game.Data) {
	e.GET("/status", getStatusHandler(channel))
	e.POST("/add", addPlayerHandler(channel))
	e.POST("/kill", killPlayerHandler(channel))
	e.POST("/move", movePlayerHandler(channel))
}
