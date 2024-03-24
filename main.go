package main

import "fmt"

func main() {
	deck := newDeck()

	deck.shuffle()

	fmt.Println(deck)
}
