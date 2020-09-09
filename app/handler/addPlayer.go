package handler

import (
	"esgi_go/game"
	p "esgi_go/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addPlayerHandler = func(channel chan game.Data) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rm *p.Player
		err := c.ShouldBindJSON(&rm)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		newPlayer := p.Player{
			Name:  rm.Name,
			Death: rm.Death,
			Kill:  rm.Kill,
		}

		channel <- game.Data{
			Message:     "addPlayer",
			PlayersData: []p.Player{newPlayer},
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusAccepted, "Player Added!")
	}
}
