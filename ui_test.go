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

func TestPromptShufflePanicsWhenNumOfCardsIsZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptShuffle(0)
}

func TestPromptShufflePanicsWhenNumOfCardsIsNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptShuffle(-1)
}

func TestPromptShuffleReturnsZeroWhenNumOfCardsIsOne(t *testing.T) {
	const EXPECTED = false
	const NUM_CARDS = 1

	actual := promptShuffle(NUM_CARDS)

	if actual != EXPECTED {
		t.Errorf("Expected %v, but got %v", EXPECTED, actual)
	}
}

func TestWillShufflePanicsWhenNumOfCardsIsZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	willShuffle(0)
}

func TestWillShufflePanicsWhenNumOfCardsIsNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	willShuffle(-1)
}

func TestWillShuffleReturnsZeroWhenNumOfCardsIsOne(t *testing.T) {
	const EXPECTED = false
	const NUM_CARDS = 1

	actual := willShuffle(NUM_CARDS)

	if actual != EXPECTED {
		t.Errorf("Expected %v, but got %v", EXPECTED, actual)
	}
}