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

func TestFormatStatusCorrectlyWhenPlayerWithFirstPersonNameHasNoCard(t *testing.T) {
	var player player
	player.name = "I"
	player.thirdPerson = false
	player.hand = []string{}
	expected := player.name + " have 0 card"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}

func TestFormatStatusCorrectlyWhenPlayerWithThirdPersonNameHasNoCard(t *testing.T) {
	var player player
	player.name = "User"
	player.thirdPerson = true
	player.hand = []string{}
	expected := player.name + " has 0 card"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}

func TestFormatStatusCorrectlyWhenPlayerWithFirstPersonNameHasOneCard(t *testing.T) {
	var player player
	player.name = "I"
	player.thirdPerson = false
	player.hand = []string{"Ace"}
	expected := player.name + " have 1 card"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}

func TestFormatStatusCorrectlyWhenPlayerWithThirdPersonNameHasOneCard(t *testing.T) {
	var player player
	player.name = "User"
	player.thirdPerson = true
	player.hand = []string{"Ace"}
	expected := player.name + " has 1 card"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}

func TestFormatStatusCorrectlyWhenPlayerWithFirstPersonNameHasTwoCards(t *testing.T) {
	var player player
	player.name = "I"
	player.thirdPerson = false
	player.hand = []string{"Ace", "Two"}
	expected := player.name + " have 2 cards"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}

func TestFormatStatusCorrectlyWhenPlayerWithThirdPersonNameHasTwoCards(t *testing.T) {
	var player player
	player.name = "User"
	player.thirdPerson = true
	player.hand = []string{"Ace", "Two"}
	expected := player.name + " has 2 cards"

	actual := formatStatus(player)

	if actual != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)		
	}
}