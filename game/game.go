package goscatgame

import (
	"container/list"
	"fmt"
	deck "goscat/deck"
	player "goscat/player"

	"github.com/google/uuid"
)

type GameState int

const (
	Creating     GameState = iota + 1
	Created                = iota + 1
	Initializing           = iota + 1
	Ready                  = iota + 1
	Connecting             = iota + 1
	Playing                = iota + 1
	Ending                 = iota + 1
	Report                 = iota + 1
	Error                  = iota + 1
)

type Game struct {

	// UUID of the game (could use for secret initially?)
	ID uuid.UUID

	// Players in the game
	Players []player.Player

	// TurnNumber that the game is currently on
	TurnNumber int

	// GameState determines what stage the game is at
	GameState GameState

	// TurnOrder for players
	TurnOrder list.List

	// CurrentPlayerTurn notes who's turn it is (used for validation)
	CurrentPlayerTurn *player.Player

	// DrawDeck used for drawing (top is concealed)
	DrawDeck deck.Deck

	// DiscardDeck used for discarding (top is revealed)
	DiscardDeck deck.Deck
}

func CreateGame(game *Game) {
	if game.GameState >= Creating {
		fmt.Printf("Game already being created!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	fmt.Printf("New game being created!\n")

	game.GameState = Creating

	game.ID = uuid.New()

	game.GameState = Created

	fmt.Printf("Game %v created!\n", game.ID)
}

func InitializeGame(game *Game, players []player.Player) {
	if game.GameState >= Initializing {
		fmt.Printf("Game already initializing!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	game.GameState = Initializing
	fmt.Printf("Game is initializing!\n")

	InitializeDecks(game)
	InitializePlayers(game, players)

	game.GameState = Ready
	fmt.Printf("Game is ready!\n")

}

func InitializeDecks(game *Game) {
	if game.GameState != Initializing {
		fmt.Printf("Game not initializing!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	fmt.Printf("Initializing decks...\n")

	if game.DiscardDeck.Name != "" {
		// DiscardDeck is already initialized, skip!
		fmt.Printf("DiscardDeck has already been initialized\n")
	} else {
		// No cards to add for discard deck!
		// Name the deck
		game.DiscardDeck.Name = "discard"
	}

	//deck.Print(&game.DiscardDeck)

	if game.DrawDeck.Name != "" {
		fmt.Printf("DrawDeck has already been initialized\n")
	} else {
		// Add cards to the deck
		game.DrawDeck.Cards = deck.DeckOfCards
		// Name the deck
		game.DrawDeck.Name = "draw"

	}

	//deck.Print(&game.DrawDeck)

	deck.Shuffle(&game.DrawDeck)

	//deck.Print(&game.DrawDeck)

}

func InitializePlayers(game *Game, players []player.Player) {

	fmt.Printf("Initializing players...\n")

	for _, s := range players {
		AddPlayer(game, &s)
	}
}

func AddPlayer(game *Game, player *player.Player) {
	game.Players = append(game.Players, *player)
}

func Deal(game *Game) {

}

func Draw(game *Game, p *player.Player, deckName string) {

	var deckChoice *deck.Deck
	switch deckName {
	case "draw":
		{
			deckChoice = &game.DrawDeck
		}
	case "discard":
		{
			deckChoice = &game.DiscardDeck
		}
	default:
		{
			fmt.Printf("Error! Wrong deck type!\n")
			return
		}
	}

	var cardPulled = deck.Draw(deckChoice)
	player.AddCardToHand(p, cardPulled)
}
