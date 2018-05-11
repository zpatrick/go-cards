package cards

import "fmt"

// A Rank represents one of the 13 ranks in a standard French playing card deck
type Rank string

const (
	Two   Rank = "Two"
	Three Rank = "Three"
	Four  Rank = "Four"
	Five  Rank = "Five"
	Six   Rank = "Six"
	Seven Rank = "Seven"
	Eight Rank = "Eight"
	Nine  Rank = "Nine"
	Ten   Rank = "Ten"
	Jack  Rank = "Jack"
	Queen Rank = "Queen"
	King  Rank = "King"
	Ace   Rank = "Ace"
)

// Ranks returns a slice containing the standard 13 ranks in order of least to greatest (with Ace high)
func Ranks() []Rank {
	return []Rank{
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
		Ace,
	}
}

// IsFace returns true if and only if r is Jack, Queen, King, or Ace
func (r Rank) IsFace() bool {
	switch r {
	case Jack, Queen, King, Ace:
		return true
	default:
		return false
	}
}

// A Suit represents one of the 4 suits in a standard French playing card deck
type Suit string

const (
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
)

// Suits returns a slice containing the standard 4 suits in alphabetical order
func Suits() []Suit {
	return []Suit{
		Clubs,
		Diamonds,
		Hearts,
		Spades,
	}
}

// A Card represents a single playing card in the standard French deck
type Card struct {
	Suit Suit
	Rank Rank
	// Value represents the card's worth.
	// This can be set by the user to help make calculating scores easier.
	// It is recommended to use a ValueMapper to set the value of each card when creating a new deck.
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

// A ValueMapper uses a card's suit and rank to calculate its value.
// This can be used to assign custom values to an entire deck of of cards when calling the NewDeck function.
type ValueMapper func(s Suit, r Rank) int

// StandardAceHigh is a ValueMapper that maps ranks with the following values (regardless of suit):
// Two:   2
// Three: 3
// Four:  4
// Five:  5
// Six:   6
// Seven: 7
// Eight: 8
// Nine:  9
// Ten:   10
// Jack:  11
// Queen: 12
// King:  13
// Ace:   14
func StandardAceHigh(s Suit, r Rank) int {
	switch r {
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 11
	case Queen:
		return 12
	case King:
		return 13
	case Ace:
		return 14
	default:
		return 0
	}
}

// StandardAceLow is a ValueMapper that maps ranks with the following values (regardless of suit):
// Ace:   1
// Two:   2
// Three: 3
// Four:  4
// Five:  5
// Six:   6
// Seven: 7
// Eight: 8
// Nine:  9
// Ten:   10
// Jack:  11
// Queen: 12
// King:  13
func StandardAceLow(s Suit, r Rank) int {
	if r == Ace {
		return 1
	}

	return StandardAceHigh(s, r)
}
