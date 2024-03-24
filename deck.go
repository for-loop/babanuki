package main

import "strings"

type deck []string

func newDeck() deck {
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	d := []string{}
	
	for i := 0; i < 4; i++ {
		d = append(d, cardValues...)
	}

	d = append(d, "Joker")

	return d
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
