package game

import (
	"fmt"
	client "goscat/pkg/client"
	"os"
	"strings"
)

// ===================================================================
//
type Player struct {

	// Name of the player
	Name string

	// Coins represent buy-ins remaining
	Coins []Coin

	// Hand is where your cards are
	Hand []Card

	// Honor represents the final life after your coins are gone
	Honor bool

	// Knocked represents the player triggering end-game
	Knocked bool

	// Client interface object for connect
	Client client.Client

	// Plays stores every play the player has made, last is 'current' move
	Plays []Play

	// PlayPhase describes what phase the gameplay is in for the player
	PlayPhase PlayPhase
}

// ===================================================================
//
func AddCardToHand(player *Player, card Card) {
	//fmt.Printf("Current hand... %q\n", player.Hand)
	player.Hand = append(player.Hand, card)
}

// ===================================================================
//
func RemoveCardFromHand(player *Player, card *Card) bool {
	removalIndex := -1

	for i, x := range player.Hand {
		if (x.Face == card.Face) && (x.Suit == card.Suit) {
			removalIndex = i
		}
	}

	fmt.Printf("Removal Index = %d\n", removalIndex)

	if removalIndex == -1 {
		fmt.Printf("Card not found in hand! %+v\n", card)
	} else {
		//not order perserving!
		player.Hand[removalIndex] = player.Hand[len(player.Hand)-1]
		player.Hand = player.Hand[:len(player.Hand)-1]
		return true
	}

	fmt.Printf("Cannot remove Card! %+v\n", card)
	return false
}

// ===================================================================
//
func DisplayHand(player *Player) {
	fmt.Printf("Player hand:\n")
	for i := range player.Hand {
		currCard := player.Hand[i]
		fmt.Printf("\t%s%s (%d)", currCard.Suit, currCard.Face, currCard.Value)
	}
	fmt.Println()
}

// ===================================================================
//
func GetPlay(player *Player) {

	fmt.Printf("\n%s's turn!", player.Name)

	fmt.Print("\n")
	PrettyPrintHand(player)
	fmt.Print("\n")

	switch player.PlayPhase {
	case DrawPhase:
		{
			fmt.Printf("Options: \"draw <deck>\" or \"knock\"!\n")
		}
	case DiscardPhase:
		{
			fmt.Printf("Options: \"discard <card>\"\n")
		}
	case EndPhase:
		{
			// D
		}
	default:
		{
			fmt.Printf("Incorrect phase! This shouldn't happen!\n%v\n\n", player.PlayPhase)
		}
	}

	retryPlay := false
	input := player.Client.GetInput()
	fmt.Printf("Input: %v\n", input)

	splits := strings.Split(input, " ")
	fmt.Printf("Splits: %v\n", splits)

	// Validate how many items we have in split
	if len(splits) != 3 {
		fmt.Println("Incorrect amount of splits! Format should be...")
		fmt.Printf("\t[draw|discard] [A|K|Q|J|10|9|...] [C|D|H|S]")
		GetPlay(player)
	}

	var currPlay Play
	switch player.PlayPhase {
	case DrawPhase:
		{
			var drawDeck string
			switch splits[0] {
			case "draw":
				{
					drawDeck = splits[1]
					currPlay.PlayOption = DrawOption
				}
			case "knock":
				{
					fmt.Printf("%s just knocked!!!\n", player.Name)
					player.Knocked = true
					currPlay.PlayOption = KnockOption
				}
			case "ff":
				{
					fmt.Printf("%s has forfeited!", player.Name)
					os.Exit(1)
				}
			default:
				{
					fmt.Printf("Incorrect draw, split[1]! Retry play!\n")
					retryPlay = true
				}
			}
			currPlay.Deck = drawDeck
		}
	case DiscardPhase:
		{
			if splits[0] != "discard" {
				fmt.Printf("Incorrect option! Retry play!\n")
				retryPlay = true
				break
			}

			// TODO - validate!
			var face = string(splits[1])
			var suit = string(splits[2])

			currPlay.Card = Card{
				Face: face,
				Suit: suit,
			}
			currPlay.PlayOption = DiscardOption

		}
	default:
		{
			fmt.Printf("Incorrect option! Retry play!\n")
			retryPlay = true
		}
	}

	if retryPlay {
		GetPlay(player)
	}

	// Finalize play!
	player.Plays = append(player.Plays, currPlay)
}

// ===================================================================
//
func PrettyPrintHand(player *Player) {
	handSize := len(player.Hand)
	for i := 0; i < handSize; i++ {
		//fmt.Printf("\t")
		fmt.Printf("┌─────┐  ")
	}
	fmt.Printf("\n")
	for j := 0; j < handSize; j++ {
		if len(player.Hand[j].Face) > 1 {
			fmt.Printf("│ %s  │  ", player.Hand[j].Face)
		} else {
			fmt.Printf("│ %s   │  ", player.Hand[j].Face)
		}

	}
	fmt.Printf("\n")
	for k := 0; k < handSize; k++ {
		fmt.Printf("│  %s │  ", SuitToSymbol(player.Hand[k]))
	}
	fmt.Printf("\n")
	for i := 0; i < handSize; i++ {
		//fmt.Printf("\t")
		fmt.Printf("└─────┘  ")
	}
}

// ===================================================================
//
func GetHandValue(player *Player) int {
	faceSums := map[string]int{
		"C": 0,
		"D": 0,
		"H": 0,
		"S": 0,
	}

	for _, x := range player.Hand {
		faceSums[x.Suit] = faceSums[x.Suit] + x.Value
	}

	maxSum := 0
	for _, val := range faceSums {
		if val > maxSum {
			maxSum = val
		}
	}

	return maxSum
}

// [wferrell - 2021/10/13]: in hindsight, should probably move this to game.go
// ===================================================================
//
type PlayPhase int

const (
	DrawPhase    PlayPhase = iota + 1
	DiscardPhase           = iota + 1
	EndPhase               = iota + 1
)

// ^^ same with this helper function
// ===================================================================
//
func PlayPhaseToString(phase PlayPhase) string {
	switch phase {
	case DrawPhase:
		{
			return "DrawPhase"
		}
	case DiscardPhase:
		{
			return "DiscardPhase"
		}
	case EndPhase:
		{
			return "EndPhase"
		}
	}

	return "UNKNOWN"
}

// ===================================================================
//
type PlayOption int

const (
	DrawOption    PlayOption = iota + 1
	KnockOption              = iota + 1
	DiscardOption            = iota + 1
)

// ===================================================================
//
type Play struct {
	PlayOption PlayOption
	Card       Card
	Deck       string
}
