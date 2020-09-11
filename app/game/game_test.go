package game

import (
	p "esgi_go/player"
	"strings"
	"testing"
)

func TestStartGame(t *testing.T) {
	channels := make(chan Data)

	go StartGame(channels)

	gameStatusChannel := make(chan string)
	channels <- Data{
		Message:     "status",
		SendChannel: gameStatusChannel,
	}

	gameStatus := <-gameStatusChannel

	if "[{bot0 0 0 {0 0}} {bot1 0 0 {0 3}} {bot2 0 0 {3 0}} {bot3 0 0 {3 3}}]" != gameStatus {
		t.Errorf("Game Start Initialization Failed : " + gameStatus)
	} else {
		t.Logf("Game Start Initialization Succeded")
	}

}

func TestAddPlayer(t *testing.T) {
	channels := make(chan Data)

	go StartGame(channels)

	newPlayer := p.Player{
		Name:  "test",
		Death: 0,
		Kill:  0,
	}

	channels <- Data{
		Message:     "addPlayer",
		PlayersData: []p.Player{newPlayer},
	}

	gameStatusChannel := make(chan string)
	channels <- Data{
		Message:     "status",
		SendChannel: gameStatusChannel,
	}

	gameStatus := <-gameStatusChannel

	if !strings.Contains(gameStatus, "test") {
		t.Errorf("Adding player Failed : " + gameStatus)
	} else {
		t.Logf("Adding player succeded")
	}
}

func TestKillPlayer(t *testing.T) {
	channels := make(chan Data)

	go StartGame(channels)

	killer := p.Player{Name: "bot1"}
	body := p.Player{Name: "bot2"}

	channels <- Data{
		Message:     "kill",
		PlayersData: []p.Player{killer, body},
	}

	gameStatusChannel := make(chan string)
	channels <- Data{
		Message:     "status",
		SendChannel: gameStatusChannel,
	}

	gameStatus := <-gameStatusChannel

	if !strings.Contains(gameStatus, "{bot1 1 0 ") || !strings.Contains(gameStatus, "{bot2 0 1 ") {
		t.Errorf("Killing player Failed : " + gameStatus)
	} else {
		t.Logf("Killing player succeded")
	}
}
