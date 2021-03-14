package goscatcard

type Card struct {

	// Face is the name of the card
	Face string

	// Suit is the class designation (e.g.: diamond, spade, etc.)
	Suit string

	// Value represents how many points the card is worth
	Value int
}
