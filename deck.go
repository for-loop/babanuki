package main

import (
	"fmt"
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

func throwPairs(d deck) deck {
	counter := make(map[string]int)

	for _, cardValue := range d {
		_, exists := counter[cardValue]
		if exists {
			counter[cardValue]++
		} else {
			counter[cardValue] = 1
		}
	}

	var newDeck deck

	for k, v := range counter {
		if v%2 == 1 {
			newDeck = append(newDeck, k)
		}
	}

	return newDeck
}

func printStatus(yourHand deck, theirHand deck) {
	fmt.Println("Com has ", len(theirHand), "cards")
	fmt.Println("You have", len(yourHand), "cards", yourHand)
}

func promptPick(d deck) int {
	var i int

	for {
		fmt.Print("Pick an opponent's card 1-", len(d), " from the left: ")
		_, err := fmt.Scanf("%v", &i)
		if err != nil {
			panic(err)
		}
		if 0 < i && i <= len(d) {
			return i
		}
		fmt.Println("Invalid input. Try again.")
	}
}
