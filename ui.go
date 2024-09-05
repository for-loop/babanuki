package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func promptPick(numHand int) int {
	var num int

	if numHand == 1 {
		return 0
	}

	for {
		fmt.Print("Pick an opponent's card 1-", numHand, " from the left: ")
		_, err := fmt.Scanf("%v", &num)
		if err != nil {
			panic(err)
		}
		if 0 < num && num <= numHand {
			return num - 1
		}
		fmt.Println("Invalid input. Try again.")
	}
}

func promptShuffle(numHand int) bool {
	var answer string

	if numHand == 1 {
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