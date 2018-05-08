package main

import (
	"bufio"
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	cards "github.com/zpatrick/go-cards"
)

func main() {
	Clear()
	tm.Println(SplashScreen)
	tm.Printf("Press enter to play...")
	tm.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for {
		Clear()
		tm.Flush()

		deck := cards.NewDeck(cards.Suits(), cards.Ranks(), BlackjackValues)
		deck.Shuffle()

		humanHand := []cards.Card{deck.Deal(), deck.Deal()}
		dealerHand := []cards.Card{deck.Deal(), deck.Deal()}

		if Score(humanHand) == 21 {
			fmt.Println("Blackjack!")
			scanner.Scan()
		}

		if Score(dealerHand) == 21 {
			fmt.Println("Blackjack!")
			scanner.Scan()
		}

		scanner.Scan()
	}
}
