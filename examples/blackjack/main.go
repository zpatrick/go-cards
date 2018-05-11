package main

import (
	"fmt"

	cards "github.com/zpatrick/go-cards"
)

func main() {
	for {
		deck := cards.NewDeck(cards.Suits(), cards.Ranks(), BlackjackValues)
		deck.Shuffle()

		humanHand := []cards.Card{deck.Deal(), deck.Deal()}
		dealerHand := []cards.Card{deck.Deal(), deck.Deal()}

		// todo: hit/stay logic here

		switch {
		case score(humanHand) > 21:
			fmt.Println("Player busts!")
		case score(dealerHand) > 21:
			fmt.Println("Dealer busts!")
		}

	}
}
