package main

import "testing"

func TestPromptPickPanicsWhenNumHandIsZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptPick(0)
}

func TestPromptPickPanicsWhenNumHandIsNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
    }()

	promptPick(-1)
}

func TestPromptPickReturnsZero(t *testing.T) {
	const EXPECTED = 0
	const NUM_HAND = 1

	actual := promptPick(NUM_HAND)

	if actual != EXPECTED {
		t.Errorf("Expected index %v, but got %v", EXPECTED, actual)
	}
}