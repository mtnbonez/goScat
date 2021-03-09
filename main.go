package main

import (
	card "goscat/card"
	coin "goscat/coin"
	game "goscat/game"
	client "goscat/networking/client"
	player "goscat/player"
)

func main() {
	// Define a game
	testGame := game.Game{}

	game.CreateGame(&testGame)

	//fmt.Println(testGame)

	// Attach two local players players
	Player1 := player.Player{
		Name:   "p1",
		Coins:  make([]coin.Coin, 0),
		Hand:   make([]card.Card, 0),
		Honor:  false,
		Client: client.LocalClient{},
	}

	Player2 := player.Player{
		Name:   "p2",
		Coins:  make([]coin.Coin, 0),
		Hand:   make([]card.Card, 0),
		Honor:  false,
		Client: client.LocalClient{},
	}

	Players := make([]player.Player, 2)
	Players[0] = Player1
	Players[1] = Player2

	game.InitializeGame(&testGame, Players)

	//fmt.Printf("%+v", testGame)

	// Start game

	//start here!!!
	game.Deal(&testGame)

	// first draws
	// choose first player (just pick first element right now)

	// Check for finish
	/*
		for {

			if testGame.GameState != game.Report {
				break
			}

			// Process input

			// Process game

			// Process status

		}
	*/

}
