package goscatgame

import (
	"fmt"
	card "goscat/card"
	deck "goscat/deck"
	player "goscat/player"

	"github.com/google/uuid"
)

// ===================================================================
//
type GameState int

const (
	Creating     GameState = iota + 1
	Created                = iota + 1
	Initializing           = iota + 1
	Ready                  = iota + 1
	Connecting             = iota + 1
	Connected              = iota + 1
	Playing                = iota + 1
	Ending                 = iota + 1
	Report                 = iota + 1
	Error                  = iota + 1
)

// ===================================================================
//
type Game struct {

	// UUID of the game (could use for secret initially?)
	ID uuid.UUID

	// Players in the game
	Players []*player.Player

	// TurnNumber that the game is currently on
	TurnNumber int

	// GameState determines what stage the game is at
	GameState GameState

	// TurnOrder for players
	TurnOrder []*player.Player

	// CurrentPlayerTurn notes who's turn it is (used for validation)
	CurrentPlayerTurn *player.Player

	// DrawDeck used for drawing (top is concealed)
	DrawDeck deck.Deck

	// DiscardDeck used for discarding (top is revealed)
	DiscardDeck deck.Deck

	// PlayerHasKnocked keeps track of who knocked
	// Might be able to take the logic out of Player if we use a pointer?
	PlayerHasKnocked bool
}

// ===================================================================
//
func CreateGame(game *Game) {
	if game.GameState == Creating {
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

// ===================================================================
//
func InitializeGame(game *Game, players []*player.Player) {
	if game.GameState == Initializing {
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

// ===================================================================
//
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

// ===================================================================
//
func InitializePlayers(game *Game, players []*player.Player) {

	fmt.Printf("Initializing players...\n")

	for i := range players {
		currPlayer := players[i]
		currPlayer.PlayPhase = player.DrawPhase
		AddPlayer(game, players[i])
	}
}

// ===================================================================
//
func AddPlayer(game *Game, player *player.Player) {
	game.Players = append(game.Players, player)
	game.TurnOrder = append(game.TurnOrder, player)
}

// ConnectPlayers attaches game clients to server
// ===================================================================
//
func ConnectPlayers(game *Game) {
	if game.GameState != Ready {
		fmt.Printf("Game not ready!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	game.GameState = Connecting

	for _, x := range game.Players {
		fmt.Printf("Connecting %+v\n", x.Name)
		if !x.Client.Connect() {
			fmt.Printf("Player %+v could not connect!\n", x.Name)
			return
		} else {
			fmt.Printf("%+v connected!\n", x.Name)
		}
	}

	game.GameState = Connected
}

// StartGame kicks the game off!
// ===================================================================
//
func StartGame(game *Game) {
	if game.GameState != Connected {
		fmt.Printf("Game players not connected!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	game.GameState = Playing
	fmt.Printf("Game now playing!\n")

	// Select first player to go first
	game.CurrentPlayerTurn = game.Players[0]
	game.TurnNumber = 1
}

// Deal out cards to players at the start of the game
// ===================================================================
//
func Deal(game *Game) {
	for i := range game.Players {
		Draw(game, game.Players[i], "draw")
		Draw(game, game.Players[i], "draw")
		Draw(game, game.Players[i], "draw")
	}

	fmt.Printf("First hands dealt!\n")
	for _, x := range game.Players {
		fmt.Printf("%q's hand: %v\n", x.Name, x.Hand)
	}

	// Discard first card from deck
	cardToDiscard := deck.Draw(&game.DrawDeck)
	deck.Discard(&game.DiscardDeck, &cardToDiscard)
}

// ===================================================================
//
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
	fmt.Printf("Card pulled: %q%q (%d)\n", cardPulled.Suit, cardPulled.Face, cardPulled.Value)
	player.AddCardToHand(p, cardPulled)
}

// ===================================================================
//
func Discard(game *Game, p *player.Player, card *card.Card) bool {
	if player.RemoveCardFromHand(p, card) {
		deck.Discard(&game.DiscardDeck, card)
		return true
	} else {
		fmt.Printf("Unable to discard card!\n")
	}

	return false
}

// Display should really go to it's own implementation
// ===================================================================
//
func Display(game *Game) {
	const border = "============================================"
	fmt.Println(border)
	fmt.Printf("ID=%s\n", game.ID.String())
	fmt.Printf("%s\tTurn %d, %s\n", game.CurrentPlayerTurn.Name, game.TurnNumber, player.PlayPhaseToString(game.CurrentPlayerTurn.PlayPhase))

	//fmt.Printf("\n")
	//player.DisplayHand(game.CurrentPlayerTurn)
	//fmt.Printf("\n")

	fmt.Printf("\tDrawDeckSize: %d\n", len(game.DrawDeck.Cards))
	fmt.Printf("\tDiscardDeckSize: %d\n", len(game.DiscardDeck.Cards))

	drawTopCard := game.DrawDeck.Cards[0]
	fmt.Printf("\tDraw Pile: %s%s (%d)\n", drawTopCard.Suit, drawTopCard.Face, drawTopCard.Value)

	var discardTopCard card.Card

	if len(game.DiscardDeck.Cards) > 0 {
		discardTopCard = game.DiscardDeck.Cards[0]

	} else {
		discardTopCard = card.Card{
			Suit:  "X",
			Face:  "X",
			Value: 0,
		}
	}
	fmt.Printf("\tDiscard Pile: %s%s (%d)\n", discardTopCard.Suit, discardTopCard.Face, discardTopCard.Value)

	fmt.Println(border)

	//Play around with stuff for the deck representations

	fmt.Printf("    ┌─────┐  ┌─────┐\n")

	//let's count the number of spaces we need
	discardSpace := false
	drawSpace := false
	if len(game.DiscardDeck.Cards) > 0 {
		if len(game.DiscardDeck.Cards[0].Face) > 1 {
			discardSpace = true
		}
	}

	if len(game.DrawDeck.Cards[0].Face) > 1 {
		drawSpace = true
	}

	// Face value
	if drawSpace {
		fmt.Printf("    │ %s  │  ", drawTopCard.Face)
	} else {
		fmt.Printf("    │ %s   │  ", drawTopCard.Face)
	}

	if discardSpace {
		fmt.Printf("│ %s  │\n", discardTopCard.Face)
	} else {
		fmt.Printf("│ %s   │\n", discardTopCard.Face)
	}

	// Suit value
	fmt.Printf("    │  %s │  │  %s │\n", card.SuitToSymbol(drawTopCard), card.SuitToSymbol(discardTopCard))

	fmt.Printf("    └─────┘  └─────┘\n")

	fmt.Printf("   Draw       Discard\n")

}

// ===================================================================
//
func Play(game *Game) {
	if game.CurrentPlayerTurn == nil {
		fmt.Printf("No player selected!\n\n%+v\n\n", &game)
		game.GameState = Error
		return
	}

	if game.CurrentPlayerTurn.PlayPhase == player.EndPhase {
		// Switch to new player
		game.TurnOrder = append(game.TurnOrder[1:], game.CurrentPlayerTurn)
		game.CurrentPlayerTurn = game.TurnOrder[0]
		game.CurrentPlayerTurn.PlayPhase = player.DrawPhase
		game.TurnNumber = game.TurnNumber + 1
		return
	}

	// Detect is knock came all the way back around!
	if game.CurrentPlayerTurn.Knocked {
		game.GameState = Ending
		return
	}

	loop := true

	for loop {
		player.GetPlay(game.CurrentPlayerTurn)

		switch game.CurrentPlayerTurn.PlayPhase {
		case player.DrawPhase:
			{
				// Need to check to see if player knocked!
				if game.CurrentPlayerTurn.Plays[len(game.CurrentPlayerTurn.Plays)-1].PlayOption == player.KnockOption {
					game.CurrentPlayerTurn.PlayPhase = player.EndPhase
				} else {
					// Else, they've drawn
					game.CurrentPlayerTurn.PlayPhase = player.DiscardPhase
					Draw(game, game.CurrentPlayerTurn, game.CurrentPlayerTurn.Plays[len(game.CurrentPlayerTurn.Plays)-1].Deck)
				}
				loop = false
			}
		case player.DiscardPhase:
			{
				if Discard(game, game.CurrentPlayerTurn, &game.CurrentPlayerTurn.Plays[len(game.CurrentPlayerTurn.Plays)-1].Card) {
					game.CurrentPlayerTurn.PlayPhase = player.EndPhase
					loop = false
				}
			}
		}
	}

}

// ===================================================================
//
func Process(game *Game, endGameCheck bool) {

	biggestHand := 0
	winningPlayer := 0
	tie := false

	// Process who is winning!
	for i, x := range game.Players {
		currHandValue := player.GetHandValue(x)
		fmt.Printf("%s has %d hand value!\n", x.Name, currHandValue)

		if i > 0 && biggestHand == currHandValue {
			tie = true
		} else if biggestHand < currHandValue {
			biggestHand = currHandValue
			winningPlayer = i
		}
	}

	// DEBUG
	fmt.Printf("Biggest hand value is %d\n", biggestHand)
	if endGameCheck {
		if tie {
			fmt.Printf("Tie!\n")
		} else {
			fmt.Printf("Player %d wins!\n", winningPlayer+1)
		}
	}
}
