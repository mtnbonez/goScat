package main

import (
	card "goscat/card"
	coin "goscat/coin"
	game "goscat/game"
	client "goscat/networking/client"
	player "goscat/player"
)

// ===================================================================
//
func main() {

	gameManager := game.GameManager{}
	game.InitManager(&gameManager)

	testGameOne := CreateTestGame("p1", "p2")
	testGameTwo := CreateTestGame("p3", "p4")

	game.AddGame(&gameManager, &testGameOne)
	game.AddGame(&gameManager, &testGameTwo)

	// Check for finish
	for {

		game.ProcessManager(&gameManager)

		// Any Scats?
	}

}

// ===================================================================
//
func CreateTestGame(p1 string, p2 string) game.Game {
	// Define a game
	testGame := game.Game{}

	game.CreateGame(&testGame)

	//fmt.Println(testGame)

	// Attach two local players players
	Player1 := player.Player{
		Name:   p1,
		Coins:  make([]coin.Coin, 0),
		Hand:   make([]card.Card, 0),
		Honor:  false,
		Client: client.NewLocalClient(),
	}

	Player2 := player.Player{
		Name:   p2,
		Coins:  make([]coin.Coin, 0),
		Hand:   make([]card.Card, 0),
		Honor:  false,
		Client: client.NewLocalClient(),
	}

	Players := make([]*player.Player, 2)
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
