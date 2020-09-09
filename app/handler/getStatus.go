package handler

import (
	"esgi_go/game"
	p "esgi_go/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

var getStatusHandler = func(channel chan game.Data) func(c *gin.Context) {
	return func(c *gin.Context) {
		sendChannel := make(chan []p.Player)
		channel <- game.Data{Message: "status", SendChannel: sendChannel}
		players := <-sendChannel
		c.JSON(http.StatusOK, players)
	}
}
