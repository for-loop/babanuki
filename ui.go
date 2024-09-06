package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func promptPick(numCards int) int {
	if numCards < 1 {
		panic("Number of cards must be a positive non-zero integer")
	}

	if numCards == 1 {
		return 0
	}

	var num int
	for {
		fmt.Print("Pick an opponent's card 1-", numCards, " from the left: ")
		_, err := fmt.Scanf("%v", &num)
		if err != nil {
			panic(err)
		}
		if 0 < num && num <= numCards {
			return num - 1
		}
		fmt.Println("Invalid input. Try again.")
	}
}

func promptShuffle(numCards int) bool {
	if numCards < 1 {
		panic("Number of cards must be a positive non-zero integer")
	}

	if numCards == 1 {
		return false
	}

	var answer string
	for {
		fmt.Print("Do you want to shuffle? (y/n) ")
		_, err := fmt.Scanf("%v", &answer)
		if err != nil {
			panic(err)
		}
		
		answerLower := strings.ToLower(answer)
		if answerLower == "y" {
			return true
		} else if answerLower == "n" {
			return false
		}
		
		fmt.Println("Invalid input. Try again.")
	}
}

func willShuffle(numCards int) bool {
	if numCards < 1 {
		panic("Number of cards must be a positive non-zero integer")
	}
	
	if numCards == 1 {
		return false
	}

	if rand.Intn(2) == 1 {
		return true
	}

	return false
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