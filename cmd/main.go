package main

import (
	"bufio"
	client "goscat/pkg/client"
	game "goscat/pkg/game"
	"os"
)

// ===================================================================
//
func main() {

	gameManager := game.GameManager{}
	game.InitManager(&gameManager)

	testGameOne := CreateTestGame("p1", "p2")
	//testGameTwo := CreateTestGame("p3", "p4")

	game.AddGame(&gameManager, &testGameOne)
	//game.AddGame(&gameManager, &testGameTwo)

	// Check for finish
	for gameManager.GetActiveAndReportingGames() {

		game.ProcessManager(&gameManager)

		// Any Scats?
	}

	// Quick STDIN to keep from closing window
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

// ===================================================================
//
func CreateTestGame(p1 string, p2 string) game.Game {
	// Define a game
	testGame := game.Game{}

	game.CreateGame(&testGame)

	//fmt.Println(testGame)

	// Attach two local players players
	Player1 := game.Player{
		Name:   p1,
		Coins:  make([]game.Coin, 0),
		Hand:   make([]game.Card, 0),
		Honor:  false,
		Client: client.NewLocalClient(),
	}

	Player2 := game.Player{
		Name:   p2,
		Coins:  make([]game.Coin, 0),
		Hand:   make([]game.Card, 0),
		Honor:  false,
		Client: client.NewLocalClient(),
	}

	Players := make([]*game.Player, 2)
	Players[0] = &Player1
	Players[1] = &Player2

	game.InitializeGame(&testGame, Players)

	//fmt.Printf("%+v", testGame)

	// Connect players
	game.ConnectPlayers(&testGame)

	// Start game

	game.Deal(&testGame)
	game.StartGame(&testGame)

	return testGame
}
