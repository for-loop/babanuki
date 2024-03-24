package main

type deck []string

func newDeck() deck {
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	d := []string{}
	
	for i := 0; i < 4; i++ {
		for _, cardValue := range cardValues {
			d = append(d, cardValue)
		}
	}

	d = append(d, "Joker")

	return d
}
