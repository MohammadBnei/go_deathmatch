package main

import (
	"esgi_go/game"
	"esgi_go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	channel := make(chan game.Data)
	go game.StartGame(channel)

	engine := gin.New()
	engine.Use(gin.Recovery())

	handler.InitializeRoutes(engine, channel)

	engine.Run(":3000")
}
