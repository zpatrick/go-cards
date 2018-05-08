package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandardDeck(t *testing.T) {
	assert.Equal(t, NewDeck(Suits(), Ranks(), StandardAceHigh), NewStandardDeck())
}

func TestDeckDeal(t *testing.T) {
	deck := NewStandardDeck()
	for card := deck.Deal(); len(deck) > 0; card = deck.Deal() {
		assert.False(t, deck.Contains(card))
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := NewStandardDeck()
	before := append(Deck{}, deck...)
	deck.Shuffle()
	after := append(Deck{}, deck...)

	// this isn't the most elegant thing in the world,
	// but it seems reasonable to say if N number of cards
	// haven't changed position, it wasn't a very good shuffle
	var unchanged int
	for i, card := range before {
		if card == after[i] {
			unchanged++
		}
	}

	if unchanged > 10 {
		t.Errorf("%d cards did not change position in the deck", unchanged)
	}
}
