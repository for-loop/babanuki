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

func TestNewDeckInitializesCardValuesWithIdenticalOrder(t *testing.T) {
	deck1String := newDeck().toString()
	deck2String := newDeck().toString()

	if deck1String != deck2String {
		t.Errorf("Expected %v, but got %v", deck1String, deck2String)
	}
}

func TestNewDeckWithShuffleRandomizesOrder(t *testing.T) {
	d1 := newDeck()
	d2 := newDeck()

	d2.shuffle()

	deck1String := d1.toString()
	deck2String := d2.toString()

	if deck1String == deck2String {
		t.Errorf("Expected %v, but got %v", deck1String, deck2String)
	}
}

func TestDealSplitsTheCardInto26And27Cards(t *testing.T) {
	d := newDeck()

	newDeck1, newDeck2 := deal(d)

	const EXPECTED_LENGTH_Deck1 = 26
	const EXPECTED_LENGTH_Deck2 = 27
	actual1 := len(newDeck1)
	actual2 := len(newDeck2)

	if actual1 != EXPECTED_LENGTH_Deck1 {
		t.Errorf("Expected %v cards to be in deck 1, but got %v", EXPECTED_LENGTH_Deck1, actual1)
	}

	if actual2 != EXPECTED_LENGTH_Deck2 {
		t.Errorf("Expected %v cards to be in deck 1, but got %v", EXPECTED_LENGTH_Deck2, actual2)
	}
}

func TestThrowPairsReturnsANewDeckWithUniqueCardValue(t *testing.T) {
	const CARD_VALUE = "Ace"
	
	d := deck{CARD_VALUE}
	newDeck := throwPairs(d)

	const EXPECTED_LENGTH = 1
	actual := len(newDeck)

	if actual != EXPECTED_LENGTH {
		t.Errorf("Expected the new deck with %v card, but got %v", EXPECTED_LENGTH, actual)
	}

	for _, cardValue := range newDeck {
		if cardValue != CARD_VALUE {
			t.Errorf("Expected the new deck containing %v, but it is absent", CARD_VALUE)
		}
	}
}

func TestThrowPairsReturnsANewDeckWithoutDuplicateCardValues(t *testing.T) {
	const CARD_VALUE = "Ace"
	
	d := deck{CARD_VALUE, CARD_VALUE}
	newDeck := throwPairs(d)

	const EXPECTED_LENGTH = 0
	actual := len(newDeck)

	if actual != EXPECTED_LENGTH {
		t.Errorf("Expected the new deck with %v card, but got %v", EXPECTED_LENGTH, actual)
	}

	for _, cardValue := range newDeck {
		if cardValue == CARD_VALUE {
			t.Errorf("Expected the new deck without %v, but it has one", CARD_VALUE)
		}
	}
}

func TestThrowPairsReturnsANewDeckWithOneOfTheTriplicatedCardValue(t *testing.T) {
	const CARD_VALUE = "Ace"
	
	d := deck{CARD_VALUE, CARD_VALUE, CARD_VALUE}
	newDeck := throwPairs(d)

	const EXPECTED_LENGTH = 1
	actual := len(newDeck)

	if actual != EXPECTED_LENGTH {
		t.Errorf("Expected the new deck with %v card, but got %v", EXPECTED_LENGTH, actual)
	}

	for _, cardValue := range newDeck {
		if cardValue != CARD_VALUE {
			t.Errorf("Expected the new deck containing %v, it is absent", CARD_VALUE)
		}
	}
}

func TestThrowPairsReturnsANewDeckWithoutQuadrupledCardValues(t *testing.T) {
	const CARD_VALUE = "Ace"
	
	d := deck{CARD_VALUE, CARD_VALUE, CARD_VALUE, CARD_VALUE}
	newDeck := throwPairs(d)

	const EXPECTED_LENGTH = 0
	actual := len(newDeck)

	if actual != EXPECTED_LENGTH {
		t.Errorf("Expected the new deck with %v card, but got %v", EXPECTED_LENGTH, actual)
	}

	for _, cardValue := range newDeck {
		if cardValue == CARD_VALUE {
			t.Errorf("Expected the new deck without %v, but it has one", CARD_VALUE)
		}
	}
}