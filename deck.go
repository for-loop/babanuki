package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	numYours := len(yourHand)
	numTheirs := len(theirHand)

	theirStatus := "Com has  " + strconv.Itoa(numTheirs) + " card"
	yourStatus := "You have " + strconv.Itoa(numYours) + " card"

	if 1 < numTheirs {
		fmt.Println(theirStatus + "s")
	} else {
		fmt.Println(theirStatus)
	}

	if 1 < numYours {
		fmt.Println(yourStatus + "s", yourHand)
	} else if numYours == 1 {
		fmt.Println(yourStatus, yourHand)
	} else {
		fmt.Println(yourStatus)
	}
}

func promptPick(d deck) int {
	var num int

	for {
		fmt.Print("Pick an opponent's card 1-", len(d), " from the left: ")
		_, err := fmt.Scanf("%v", &num)
		if err != nil {
			panic(err)
		}
		if 0 < num && num <= len(d) {
			return num
		}
		fmt.Println("Invalid input. Try again.")
	}
}

func takeOut(i int, d deck) (string, deck) {
	if i < 0 && len(d) <= i {
		panic("Index out of range error")
	}

	cardValue := d[i]

	return cardValue, append(d[:i], d[i+1:]...)
}

func (d deck) contains(cardValue string) bool {
	for _, c := range d {
		if c == cardValue {
			return true
		}
	}
	return false
}

func remove(cardValue string, d deck) deck {
	for i, c := range d {
		if c == cardValue {
			return append(d[:i], d[i+1:]...)
		}
	}
	return d
}

func winnerExists(yourHand deck, theirHand deck) bool {
	numYours := len(yourHand)
	numTheirs := len(theirHand)

	if numYours == 0 && numTheirs == 0 {
		panic("There cannot be more than one winner!")
	}

	if numYours == 0 {
		fmt.Println("You won! 🎉")
		return true
	}
	if numTheirs == 0 {
		fmt.Println("Com won!")
		return true
	}
	return false
}
