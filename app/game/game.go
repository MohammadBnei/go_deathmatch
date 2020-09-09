package game

import (
	p "esgi_go/player"
	"strconv"
)

type Game struct {
	Players []p.Player
}

type Data struct {
	Message     string
	PlayersData []p.Player
	GameData    Game
	SendChannel chan []p.Player
}

func StartGame(channel chan Data) {
	bots := []p.Player{}
	for i := 0; i < 4; i++ {
		b := p.Player{Name: "bot" + strconv.Itoa(i), Death: 0, Kill: 0}
		bots = append(bots, b)
	}

	game := Game{Players: bots}

	for {
		data := <-channel
		switch data.Message {
		case "addPlayer":
			addPlayer(&game, data)
		case "kill":
			killPlayer(&game, data)
		case "status":
			go sendData(&game, data.SendChannel)
		}
	}
}

func addPlayer(game *Game, data Data) {
	newPlayer := data.PlayersData[0]
	game.Players = append(game.Players, newPlayer)
}

func killPlayer(game *Game, data Data) {
	killer := data.PlayersData[0]
	body := data.PlayersData[1]

	for i := range game.Players {
		switch game.Players[i].Name {
		case killer.Name:
			game.Players[i].Kill++
			break
		case body.Name:
			game.Players[i].Death++
			break
		}
	}
}

func sendData(game *Game, channel chan []p.Player) {
	channel <- game.Players
}
