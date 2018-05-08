package main

import cards "github.com/zpatrick/go-cards"

func BlackjackValues(s cards.Suit, r cards.Rank) int {
	switch r {
	case cards.Ace:
		return 1
	case cards.Two:
		return 2
	case cards.Three:
		return 3
	case cards.Four:
		return 4
	case cards.Five:
		return 5
	case cards.Six:
		return 6
	case cards.Seven:
		return 7
	case cards.Eight:
		return 8
	case cards.Nine:
		return 9
	case cards.Ten, cards.Jack, cards.Queen, cards.King:
		return 10
	default:
		return 0
	}
}

func Score(hand []cards.Card) int {
	return 0
}
