package cards

import (
	"math/rand"
)

type Deck []Card

func NewDeck(suits []Suit, ranks []Rank, valueMapper ValueMapper) Deck {
	deck := make(Deck, 0, len(suits)*len(ranks))
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

func NewStandardDeck() Deck {
	return NewDeck(Suits(), Ranks(), StandardAceHigh)
}

func (d Deck) Contains(c Card) bool {
	return d.ContainsFunc(func(card Card) bool {
		return card == c
	})
}

func (d Deck) ContainsFunc(f func(c Card) bool) bool {
	for _, c := range d {
		if f(c) {
			return true
		}
	}

	return false
}

func (d *Deck) Deal() Card {
	card := (*d)[len(*d)-1]
	(*d) = (*d)[:len(*d)-1]
	return card
}

func (d Deck) Shuffle() {
	for i, j := range rand.Perm(len(d)) {
		d[i], d[j] = d[j], d[i]
	}
}
