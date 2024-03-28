package main

import "fmt"

func main() {
	deck := newDeck()

	deck.shuffle()

	yourHand, theirHand := deal(deck)

	yourHand = throwPairs(yourHand)
	theirHand = throwPairs(theirHand)

	printStatus(yourHand, theirHand)
	num := promptPick(theirHand)
	i := num - 1

	fmt.Println("You picked", theirHand[i])

	cardValue, _ := takeOut(i, theirHand)
	fmt.Println("Pulled out", cardValue)
}
