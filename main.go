package main

import "fmt"

func main() {
	deck := newDeck()

	deck.shuffle()

	yourHand, theirHand := deal(deck)

	yourHand = throwPairs(yourHand)
	theirHand = throwPairs(theirHand)

	printStatus(yourHand, theirHand)
	i := promptPick(theirHand)

	fmt.Println("You picked", theirHand[i-1])
}
