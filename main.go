package main

import "fmt"

func main() {
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	deck := []string{}
	
	for i := 0; i < 4; i++ {
		for _, cardValue := range cardValues {
			deck = append(deck, cardValue)
		}
	}

	deck = append(deck, "Joker")

	fmt.Println(deck)
}
