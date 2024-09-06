package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := newDeck()

	deck.shuffle()

	your, their := deal(deck)

	your.hand = throwPairs(your.hand)
	their.hand = throwPairs(their.hand)

	printStatus(your.hand, their.hand)

	cardValue := ""

	for {
		i := promptPick(len(their.hand))

		cardValue, their.hand = takeOut(i, their.hand)
		fmt.Println("You picked Com's", cardValue)

		if your.hand.contains(cardValue) {
			fmt.Println("You already have", cardValue)
			your.hand = remove(cardValue, your.hand)
			fmt.Println("You throw out", cardValue)
		} else {
			fmt.Println("You add", cardValue)
			your.hand = append(your.hand, cardValue)
		}
		printStatus(your.hand, their.hand)

		if winnerExists(your.hand, their.hand) {
			break
		}

		if promptShuffle(len(your.hand)) {
			your.hand.shuffle()
			fmt.Println("You shuffled your hand")
			printStatus(your.hand, their.hand)
		}

		time.Sleep(1 * time.Second)

		i = rand.Intn(len(your.hand))

		cardValue, your.hand = takeOut(i, your.hand)
		fmt.Println("Com took your", cardValue)

		if their.hand.contains(cardValue) {
			fmt.Println("Com already has", cardValue)
			their.hand = remove(cardValue, their.hand)
			fmt.Println("Com throws out", cardValue)
		} else {
			fmt.Println("Com adds", cardValue)
			their.hand = append(their.hand, cardValue)
		}
		printStatus(your.hand, their.hand)

		if winnerExists(your.hand, their.hand) {
			break
		}

		if willShuffle(len(their.hand)) {
			their.hand.shuffle()
			fmt.Println("Com shuffled its hand")
			printStatus(your.hand, their.hand)
		}
	}
}
