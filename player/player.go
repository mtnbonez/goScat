package goscatplayer

import (
	"fmt"
	card "goscat/card"
	coin "goscat/coin"
	client "goscat/networking/client"
	"strings"
)

// ===================================================================
//
type Player struct {

	// Name of the player
	Name string

	// Coins represent buy-ins remaining
	Coins []coin.Coin

	// Hand is where your cards are
	Hand []card.Card

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
func AddCardToHand(player *Player, card card.Card) {
	//fmt.Printf("Current hand... %q\n", player.Hand)
	player.Hand = append(player.Hand, card)
}

// ===================================================================
//
func DisplayHand(player *Player) {
	fmt.Printf("Player hand:\n")
	for i := range player.Hand {
		currCard := player.Hand[i]
		fmt.Printf("\t%q%q (%d)", currCard.Suit, currCard.Face, currCard.Value)
	}
	fmt.Println()
}

// ===================================================================
//
func GetPlay(player *Player) {

	DisplayHand(player)

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
	splits := strings.Split(input, " ")

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Splits: %v\n", splits)

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
					fmt.Printf("%q just knocked!!!\n", player.Name)
					player.Knocked = true
					currPlay.PlayOption = KnockOption
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
			var suit = string(splits[1])
			var face = string(splits[2])

			currPlay.Card = card.Card{
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

// ===================================================================
//
type PlayPhase int

const (
	DrawPhase    PlayPhase = iota + 1
	DiscardPhase           = iota + 1
	EndPhase               = iota + 1
)

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
	Card       card.Card
	Deck       string
}
