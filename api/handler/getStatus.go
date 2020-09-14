package handler

import (
	"esgi_go/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

var getStatusHandler = func(channel chan game.Data) func(c *gin.Context) {
	return func(c *gin.Context) {
		sendChannel := make(chan string)
		channel <- game.Data{Message: "status", SendChannel: sendChannel}
		players := <-sendChannel
		c.JSON(http.StatusOK, players)
	}
}
