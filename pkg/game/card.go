package game

type Card struct {

	// Suit is the class designation (e.g.: diamond, spade, etc.)
	Suit string

	// Face is the name of the card
	Face string

	// Value represents how many points the card is worth
	Value int
}

func SuitToSymbol(card Card) string {
	switch card.Suit {
	case "S":
		{
			return "♠️"
		}
	case "H":
		{
			return "♥️"
		}
	case "C":
		{
			return "♣️"
		}
	case "D":
		{
			return "♦️"
		}
	}

	return "?"
}
