package holdem

import (
	"testing"
)

func TestHand_isRoyalFlush(t *testing.T) {
	hand := Hand{
		Name: "Test Hand",
		Cards: []Card{
			{Suit: "S", Name: "T"},
			{Suit: "S", Name: "J"},
			{Suit: "S", Name: "Q"},
			{Suit: "S", Name: "K"},
			{Suit: "S", Name: "A"},
		},
	}

	result := hand.isRoyalFlush()
	if result == nil {
		t.Error("Expected Royal Flush, but got nil")
	} else if result.CombinationName != "Royal Flush" {
		t.Errorf("Expected Royal Flush, but got %s", result.CombinationName)
	}
}

func TestHand_isStraightFlush(t *testing.T) {
	// Test case 1: Straight Flush
	hand1 := Hand{
		Name: "Straight Flush",
		Cards: []Card{
			{Suit: "S", Name: "7", Weight: 7},
			{Suit: "S", Name: "8", Weight: 8},
			{Suit: "S", Name: "9", Weight: 9},
			{Suit: "S", Name: "T", Weight: 10},
			{Suit: "S", Name: "J", Weight: 11},
		},
	}
	result1 := hand1.isStraightFlush()
	if result1 == nil || result1.CombinationName != "Straight Flush" {
		t.Errorf("Test case 1 failed: Expected Straight Flush")
	}

	// Test case 2: Non-Straight Flush
	hand2 := Hand{
		Name: "Non-Straight Flush Hand",
		Cards: []Card{
			{Suit: "S", Name: "A"},
			{Suit: "H", Name: "2"},
			{Suit: "C", Name: "3"},
			{Suit: "D", Name: "4"},
			{Suit: "H", Name: "5"},
		},
	}
	result2 := hand2.isStraightFlush()
	if result2 != nil {
		t.Errorf("Test case 2 failed: Expected nil, but got %s", result2.CombinationName)
	}
}

func TestHand_isFullHouse(t *testing.T) {
	// Implement test cases for isFullHouse
}

func TestHand_isFourOfAKind(t *testing.T) {
	// Implement test cases for isFourOfAKind
}

func TestHand_isFlush(t *testing.T) {
	// Implement test cases for isFlush
}

func TestHand_isStraight(t *testing.T) {
	// Implement test cases for isStraight
}

func TestHand_isThreeOfAKind(t *testing.T) {
	// Implement test cases for isThreeOfAKind
}

func TestHand_isTwoPair(t *testing.T) {
	// Implement test cases for isTwoPair
}

func TestHand_isPair(t *testing.T) {
	// Implement test cases for isPair
}

func TestHand_isHighCard(t *testing.T) {
	// Implement test cases for isHighCard
}

func TestHand_DefineCombination(t *testing.T) {
	// Implement test cases for DefineCombination
}
