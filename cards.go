package cards

import (
	"math/rand"
	"sort"
)

// Cards is a named type for []Card.
// This type can be used to represent a deck or hand of cards.
type Cards []Card

// NewDeck creates a slice of Cards containing one card for each suit-rank pair.
// The valueMapper is used to assign a value to each card in the slice.
func NewDeck(suits []Suit, ranks []Rank, valueMapper ValueMapper) Cards {
	deck := make(Cards, 0, len(suits)*len(ranks))
	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{
				Suit:  suit,
				Rank:  rank,
				Value: valueMapper(suit, rank),
			}

			deck = append(deck, card)
		}
	}

	return deck
}

// NewStandardDeck returns a standard 52-card deck with Ace-high values.
// This is shorthand for:
//  NewDeck(Suits(), Ranks(), StandardAceHigh)
func NewStandardDeck() Cards {
	return NewDeck(Suits(), Ranks(), StandardAceHigh)
}

// Contains returns true if card is within c.
func (c Cards) Contains(card Card) bool {
	return c.ContainsFunc(func(other Card) bool {
		return card == other
	})
}

// ContainsFunc returns true if any card in c satisfies f(card).
func (c Cards) ContainsFunc(f func(card Card) bool) bool {
	for _, card := range c {
		if f(card) {
			return true
		}
	}

	return false
}

// Deal will remove the last card from c and return it.
// A zero-value Card will be returned if len(c) == 0.
func (c *Cards) Deal() Card {
	if c == nil || len(*c) == 0 {
		return Card{}
	}

	card := (*c)[len(*c)-1]
	(*c) = (*c)[:len(*c)-1]
	return card
}

// Shuffle will randomly change the indexes of cards in c.
func (c Cards) Shuffle() {
	for i, j := range rand.Perm(len(c)) {
		c[i], c[j] = c[j], c[i]
	}
}

// Len reports the number of elements in c.
func (c Cards) Len() int {
	return len(c)
}

// Less returns true if and only if the value the card at index i
// is less than the value of the card at index j
func (c Cards) Less(i, j int) bool {
	return c[i].Value < c[j].Value
}

// Swap swaps cards with indexes i and j.
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Sort sorts the cards by their value.
// This is shorthand for:
//  sort.Sort(c)
func (c Cards) Sort() {
	sort.Sort(c)
}
