package goscatdeck

import (
	"fmt"
	card "goscat/card"
	"math/rand"
	"time"
)

// Deck is used for deques of cards
type Deck struct {

	// Name of the deck (e.g.: discard, draw, etc.)
	Name string

	// Cards
	Cards []card.Card
}

func Draw(d *Deck) card.Card {
	top := d.Cards[0]
	d.Cards = d.Cards[1:]
	return top
}

func Discard(d *Deck, c *card.Card) {
	// Prepend for the discard pile! [0] is always visible
	d.Cards = append([]card.Card{*c}, d.Cards...)
}

func Shuffle(d *Deck) {
	rand.Seed(time.Now().Unix())
	dest := make([]card.Card, len(d.Cards))
	perm := rand.Perm(len(d.Cards))
	for i, v := range perm {
		dest[v] = d.Cards[i]
	}

	d.Cards = dest
}

func Print(d *Deck) {
	fmt.Printf("%q\n", d.Name)
	for i, v := range d.Cards {
		fmt.Printf("[%d] %q%c\n", i, v.Face, v.Suit)
	}
}

// "◆" "♥" "♣" "♠"
var DeckOfCards = []card.Card{
	{Face: "2", Suit: "D", Value: 2},
	{Face: "3", Suit: "D", Value: 3},
	{Face: "4", Suit: "D", Value: 4},
	{Face: "5", Suit: "D", Value: 5},
	{Face: "6", Suit: "D", Value: 6},
	{Face: "7", Suit: "D", Value: 7},
	{Face: "8", Suit: "D", Value: 8},
	{Face: "9", Suit: "D", Value: 9},
	{Face: "10", Suit: "D", Value: 10},
	{Face: "J", Suit: "D", Value: 10},
	{Face: "Q", Suit: "D", Value: 10},
	{Face: "K", Suit: "D", Value: 10},
	{Face: "A", Suit: "D", Value: 11},
	{Face: "2", Suit: "H", Value: 2},
	{Face: "3", Suit: "H", Value: 3},
	{Face: "4", Suit: "H", Value: 4},
	{Face: "5", Suit: "H", Value: 5},
	{Face: "6", Suit: "H", Value: 6},
	{Face: "7", Suit: "H", Value: 7},
	{Face: "8", Suit: "H", Value: 8},
	{Face: "9", Suit: "H", Value: 9},
	{Face: "10", Suit: "H", Value: 10},
	{Face: "J", Suit: "H", Value: 10},
	{Face: "Q", Suit: "H", Value: 10},
	{Face: "K", Suit: "H", Value: 10},
	{Face: "A", Suit: "H", Value: 11},
	{Face: "2", Suit: "S", Value: 2},
	{Face: "3", Suit: "S", Value: 3},
	{Face: "4", Suit: "S", Value: 4},
	{Face: "5", Suit: "S", Value: 5},
	{Face: "6", Suit: "S", Value: 6},
	{Face: "7", Suit: "S", Value: 7},
	{Face: "8", Suit: "S", Value: 8},
	{Face: "9", Suit: "S", Value: 9},
	{Face: "10", Suit: "S", Value: 10},
	{Face: "J", Suit: "S", Value: 10},
	{Face: "Q", Suit: "S", Value: 10},
	{Face: "K", Suit: "S", Value: 10},
	{Face: "A", Suit: "S", Value: 11},
	{Face: "2", Suit: "C", Value: 2},
	{Face: "3", Suit: "C", Value: 3},
	{Face: "4", Suit: "C", Value: 4},
	{Face: "5", Suit: "C", Value: 5},
	{Face: "6", Suit: "C", Value: 6},
	{Face: "7", Suit: "C", Value: 7},
	{Face: "8", Suit: "C", Value: 8},
	{Face: "9", Suit: "C", Value: 9},
	{Face: "10", Suit: "C", Value: 10},
	{Face: "J", Suit: "C", Value: 10},
	{Face: "Q", Suit: "C", Value: 10},
	{Face: "K", Suit: "C", Value: 10},
	{Face: "A", Suit: "C", Value: 11},
}
