package cards

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandardDeck(t *testing.T) {
	assert.Equal(t, NewDeck(Suits(), Ranks(), StandardAceHigh), NewStandardDeck())
}

func TestCardsDeal(t *testing.T) {
	deck := NewStandardDeck()
	for card := deck.Deal(); len(deck) > 0; card = deck.Deal() {
		assert.False(t, deck.Contains(card))
	}

	// dealing on an empty deck should return a zero-value card
	assert.Equal(t, Card{}, deck.Deal())
}

func TestCardsShuffle(t *testing.T) {
	deck := NewStandardDeck()
	before := append(Cards{}, deck...)
	deck.Shuffle()
	after := append(Cards{}, deck...)

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

func TestCardsSort(t *testing.T) {
	deck := NewStandardDeck()
	deck.Shuffle()
	deck.Sort()

	for i := 0; i < len(deck)-2; i++ {
		if deck[i].Value > deck[i+1].Value {
			text := fmt.Sprintf("card at index %d (%#v) ", i, deck[i])
			text += fmt.Sprintf("has a greater value than ")
			text += fmt.Sprintf("card at index %d (%#v) ", i+1, deck[i+1])
			t.Fatal(text)
		}
	}
}
