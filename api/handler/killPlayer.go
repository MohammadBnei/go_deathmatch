package handler

import (
	"esgi_go/game"
	p "esgi_go/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kill struct {
	Killer string `json:"killer"`
	Body   string `json:"body"`
}

var killPlayerHandler = func(channel chan game.Data) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rm *kill
		err := c.ShouldBindJSON(&rm)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		killer := p.Player{Name: rm.Killer}
		body := p.Player{Name: rm.Body}

		channel <- game.Data{
			Message:     "kill",
			PlayersData: []p.Player{killer, body},
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusAccepted, killer.Name+" killed "+body.Name)
	}
}
