package main

import "fmt"

func main() {
	deck := newDeck()

	deck.shuffle()

	yourHand, _ := deal(deck)

	yourHand = throwPairs(yourHand)

	fmt.Println(yourHand)
}
