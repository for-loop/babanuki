package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := newDeck()

	deck.shuffle()

	yourHand, theirHand := deal(deck)

	yourHand = throwPairs(yourHand)
	theirHand = throwPairs(theirHand)

	printStatus(yourHand, theirHand)

	cardValue := ""

	for {
		num := promptPick(theirHand)
		i := num - 1

		cardValue, theirHand = takeOut(i, theirHand)
		fmt.Println("You picked Com's", cardValue)

		if yourHand.contains(cardValue) {
			fmt.Println("You already have", cardValue)
			yourHand = remove(cardValue, yourHand)
			fmt.Println("You throw out", cardValue)
		} else {
			fmt.Println("You add", cardValue)
			yourHand = append(yourHand, cardValue)
		}
		printStatus(yourHand, theirHand)

		if winnerExists(yourHand, theirHand) {
			break
		}

		time.Sleep(1 * time.Second)

		i = rand.Intn(len(yourHand))

		cardValue, yourHand = takeOut(i, yourHand)
		fmt.Println("Com took your", cardValue)

		if theirHand.contains(cardValue) {
			fmt.Println("Com already has", cardValue)
			theirHand = remove(cardValue, theirHand)
			fmt.Println("Com throws out", cardValue)
		} else {
			fmt.Println("Com adds", cardValue)
			theirHand = append(theirHand, cardValue)
		}
		printStatus(yourHand, theirHand)

		if winnerExists(yourHand, theirHand) {
			break
		}
	}
}
