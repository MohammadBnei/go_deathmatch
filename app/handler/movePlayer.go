package handler

import (
	"esgi_go/game"
	p "esgi_go/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

type move struct {
	Player string
	Move   string
}

var movePlayerHandler = func(channel chan game.Data) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rm *move
		err := c.ShouldBindJSON(&rm)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		channel <- game.Data{
			Message:     "move",
			PlayersData: []p.Player{p.Player{Name: rm.Player}},
			Move:        rm.Move,
		}

		if err != nil {
			panic(err)
		}

		movement := move{
			Player: rm.Player,
			Move:   rm.Move,
		}

		c.JSON(http.StatusAccepted, movement.Player+" moved "+movement.Move)
	}
}
