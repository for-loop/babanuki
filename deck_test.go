package main

import "testing"

func TestNewDeckContains53Cards(t *testing.T) {
	d := newDeck()

	const EXPECTED = 53
	actual := len(d)

	if actual != EXPECTED {
		t.Errorf("Expected %v cards in the new deck, but got %v", EXPECTED, actual)
	}
}

func TestNewDeckHasAceAsTheFirstCardValue(t *testing.T) {
	d := newDeck()

	const EXPECTED = "Ace"
	actual := d[0]

	if actual != EXPECTED {
		t.Errorf("Expected %v as the first card value, but got %v", EXPECTED, actual)
	}
}

func TestNewDeckHasJokerAsTheLastCardValue(t *testing.T) {
	d := newDeck()

	const EXPECTED = "Joker"
	actual := d[len(d) - 1]

	if actual != EXPECTED {
		t.Errorf("Expected %v as the first card value, but got %v", EXPECTED, actual)
	}
}