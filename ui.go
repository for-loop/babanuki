package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func promptShuffle(d deck) bool {
	var answer string

	if len(d) == 1 {
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