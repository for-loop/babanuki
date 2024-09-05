package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func promptPick(numCards int) int {
	if numCards < 1 {
		panic("Number of hand must be a positive non-zero integer")
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
	var answer string

	if numCards == 1 {
		return false
	}

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

func willShuffle(d deck) bool {
	if len(d) == 1 {
		return false
	}

	if rand.Intn(2) == 1 {
		return true
	}

	return false
}