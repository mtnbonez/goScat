package goscatplayer

import (
	"fmt"
	card "goscat/card"
	coin "goscat/coin"
	client "goscat/networking/client"
	"strings"
)

type Player struct {

	// Name of the player
	Name string

	// Coins represent buy-ins remaining
	Coins []coin.Coin

	// Hand is where your cards are
	Hand []card.Card

	// Honor represents the final life after your coins are gone
	Honor bool

	// Client interface object for connect
	Client client.Client

	// Plays stores every play the player has made, last is 'current' move
	Plays []Play

	// PlayPhase describes what phase the gameplay is in for the player
	PlayPhase PlayPhase
}

func AddCardToHand(player *Player, card card.Card) {
	player.Hand = append(player.Hand, card)
}

func GetPlay(player *Player) {
	retryPlay := false
	input := player.Client.GetInput()
	splits := strings.Split(input, " ")

	var currPlay Play
	switch player.PlayPhase {
	case DrawPhase:
		{
			var drawDeck string
			switch splits[0] {
			case "draw":
				{
					drawDeck = splits[1]
				}
			default:
				{
					fmt.Printf("Incorrect option! Retry play!\n")
					retryPlay = true
				}
			}
			currPlay.Deck = drawDeck
			currPlay.PlayOption = DrawOption
		}
	case DiscardPhase:
		{
			if splits[0] != "discard" {
				fmt.Printf("Incorrect option! Retry play!\n")
				retryPlay = true
				break
			}

			// TODO - validate!
			var cardName = splits[1]
			var suit = rune(cardName[0])
			var face = string(cardName[1])

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

type PlayPhase int

const (
	DrawPhase    PlayPhase = iota + 1
	DiscardPhase           = iota + 1
	EndPhase               = iota + 1
)

type PlayOption int

const (
	DrawOption    PlayOption = iota + 1
	DiscardOption            = iota + 1
)

type Play struct {
	PlayOption PlayOption
	Card       card.Card
	Deck       string
}
