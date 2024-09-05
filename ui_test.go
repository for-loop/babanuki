package main

import "testing"

func TestPromptPickPanicsWhenNumOfCardsIsZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptPick(0)
}

func TestPromptPickPanicsWhenNumOfCardsIsNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptPick(-1)
}

func TestPromptPickReturnsZeroWhenNumOfCardsIsOne(t *testing.T) {
	const EXPECTED = 0
	const NUM_CARDS = 1

	actual := promptPick(NUM_CARDS)

	if actual != EXPECTED {
		t.Errorf("Expected index %v, but got %v", EXPECTED, actual)
	}
}