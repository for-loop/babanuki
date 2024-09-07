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

func printStatus(your player, their player) {
	theirStatus := formatStatus(their)
	yourStatus := formatStatus(your)

	fmt.Println(theirStatus)

	if len(your.hand) == 0 {
		fmt.Println(yourStatus)
	} else {
		fmt.Println(yourStatus, your.hand)
	}
}

func formatStatus(player player) string {
	status := player.name

	if player.thirdPerson {
		status += " has "
	} else {
		status += " have "
	}

	numCards := len(player.hand)
	status += strconv.Itoa(numCards) + " card"

	if 1 < numCards {
		status += "s"
	}

	return status
}
