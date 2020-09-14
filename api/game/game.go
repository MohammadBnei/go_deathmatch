package game

import (
	"errors"
	p "esgi_go/player"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Players [4]p.Player
}

type Data struct {
	Message     string
	Move        string
	PlayersData []p.Player
	GameData    Game
	SendChannel chan string
}

func StartGame(channel chan Data) {
	bots := [4]p.Player{}
	for i := 0; i < 4; i++ {
		var position p.Position
		switch i {
		case 0:
			position = p.Position{X: 0, Y: 0}
		case 1:
			position = p.Position{X: 0, Y: 3}
		case 2:
			position = p.Position{X: 3, Y: 0}
		case 3:
			position = p.Position{X: 3, Y: 3}

		}
		b := p.Player{
			Name:     "bot" + strconv.Itoa(i),
			Death:    0,
			Kill:     0,
			Position: position,
		}

		bots[i] = b
	}

	game := Game{Players: bots}

	for {
		data := <-channel
		switch data.Message {
		case "addPlayer":
			addPlayer(&game, data)
		case "move":
			movePlayer(&game, data)
		case "kill":
			killPlayer(&game, data)
		case "status":
			go sendData(&game, data.SendChannel)
		}
	}
}

func addPlayer(game *Game, data Data) {
	newPlayer := data.PlayersData[0]
	for i := 0; i < 4; i++ {
		if strings.HasPrefix(game.Players[i].Name, "bot") {
			randomPosition(game, &newPlayer)
			game.Players[i] = newPlayer
			return
		}
	}
}

func killPlayer(game *Game, data Data) error {
	killer, err := getGamePlayer(game, data.PlayersData[0].Name)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	body, err := getGamePlayer(game, data.PlayersData[1].Name)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	killer.Kill++
	body.Death++

	return nil
}

func getGamePlayer(game *Game, playerName string) (*p.Player, error) {
	var player *p.Player
	for i := range game.Players {
		if game.Players[i].Name == playerName {
			player = &game.Players[i]
		}
	}

	if player == nil {
		return nil, errors.New(playerName + " not Found")
	}
	return player, nil
}

func movePlayer(game *Game, data Data) error {
	player, err := getGamePlayer(game, data.PlayersData[0].Name)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	switch data.Move {
	case "down":
		player.Position.X = ((player.Position.X + 1) % 4)
	case "up":
		if player.Position.X == 0 {
			player.Position.X = 3
			break
		}
		player.Position.X = ((player.Position.X - 1) % 4)
	case "left":
		if player.Position.Y == 0 {
			player.Position.Y = 3
			break
		}
		player.Position.Y = ((player.Position.Y - 1) % 4)
	case "right":
		player.Position.Y = ((player.Position.Y + 1) % 4)
	}

	moveKill(game, player)

	return nil
}

func moveKill(game *Game, killer *p.Player) {
	for i := range game.Players {
		if game.Players[i].Name == killer.Name {
			continue
		}
		if verifyPosition(game.Players[i], *killer) {
			killer.Kill++
			game.Players[i].Death++
			randomPosition(game, &game.Players[i])
			return
		}
	}
}

func randomPosition(game *Game, body *p.Player) {
	for i := 0; i < 4; i++ {
		var position p.Position
		switch i {
		case 0:
			position = p.Position{X: 0, Y: 0}
		case 1:
			position = p.Position{X: 0, Y: 3}
		case 2:
			position = p.Position{X: 3, Y: 0}
		case 3:
			position = p.Position{X: 3, Y: 3}
		}

		if verifyAllPosition(game, p.Player{Position: position}) {
			body.Position = position
			return
		}
	}

	body.Position = p.Position{X: 2, Y: 2}
}

func verifyPosition(player1 p.Player, player2 p.Player) bool {
	if player1.Position.X == player2.Position.X && player1.Position.Y == player2.Position.Y {
		return true
	}

	return false
}

func verifyAllPosition(game *Game, player p.Player) bool {
	for _, pl := range game.Players {
		if verifyPosition(pl, player) {
			return false
		}
	}
	return true
}

func sendData(game *Game, channel chan string) {
	gameMap := [4][4]string{}
	for x := range gameMap {
		for y := range gameMap[x] {
			gameMap[x][y] = "."
		}
	}

	for _, player := range game.Players {
		gameMap[player.Position.X][player.Position.Y] = player.Name
	}
	channel <- fmt.Sprint(game.Players)
}
