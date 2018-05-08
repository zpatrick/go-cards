package cards

import "fmt"

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

func (r Rank) IsFace() bool {
	switch r {
	case Jack, Queen, King, Ace:
		return true
	default:
		return false
	}
}

type Suit string

const (
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
)

func Suits() []Suit {
	return []Suit{
		Clubs,
		Diamonds,
		Hearts,
		Spades,
	}
}

type Card struct {
	Suit  Suit
	Rank  Rank
	Value int
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

type ValueMapper func(s Suit, r Rank) int

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

func StandardAceLow(s Suit, r Rank) int {
	if r == Ace {
		return 1
	}

	return StandardAceHigh(s, r)
}
