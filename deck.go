package main

import (
	"math/rand"
	"strings"
)

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

func (d deck) shuffle() {
	for i := range d {
		j := rand.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

func deal(d deck) (deck, deck) {
	return d[:26], d[26:]
}