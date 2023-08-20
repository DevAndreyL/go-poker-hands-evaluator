package holdem

import (
	"testing"
)

func TestHand_isRoyalFlush(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Royal Flush",
			cards: []Card{
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
				{Suit: "H", Name: "A", Weight: 14},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Royal Flush",
				HandWeight:        60,
				CombinationWeight: royalFlushCombinationWeight,
			},
		},
		{
			name: "No Royal Flush (Different Suit)",
			cards: []Card{
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "C", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
				{Suit: "H", Name: "A", Weight: 14},
			},
			expected: nil,
		},
		{
			name: "No Royal Flush (Missing Card)",
			cards: []Card{
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isRoyalFlush(), test)
	}
}

func TestHand_isStraightFlush(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Straight Flush",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Straight Flush",
				HandWeight:        55,
				CombinationWeight: straightFlushCombinationWeight,
			},
		},
		{
			name: "No Straight Flush (Different Suit)",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
			},
			expected: nil,
		},
		{
			name: "No Straight Flush (Not Enough Cards)",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isStraightFlush(), test)
	}
}

func TestHand_isFullHouse(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Full House",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "A", Weight: 14},
				{Suit: "D", Name: "A", Weight: 14},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "5", Weight: 5},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Full House",
				HandWeight:        52,
				CombinationWeight: 7,
			},
		},
		{
			name: "No Full House",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "K", Weight: 13},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isFullHouse(), test)
	}
}

func TestHand_isFourOfAKind(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Four of a Kind",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "A", Weight: 14},
				{Suit: "D", Name: "A", Weight: 14},
				{Suit: "S", Name: "A", Weight: 14},
				{Suit: "H", Name: "5", Weight: 5},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Four of a kind",
				HandWeight:        61,
				CombinationWeight: fourOfAKindCombinationWeight,
			},
		},
		{
			name: "No Four of a Kind",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "K", Weight: 13},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isFourOfAKind(), test)
	}
}

func TestHand_isFlush(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Flush",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "H", Name: "K", Weight: 13},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "5", Weight: 5},
				{Suit: "H", Name: "2", Weight: 2},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Flush",
				HandWeight:        46,
				CombinationWeight: flushCombinationWeight,
			},
		},
		{
			name: "No Flush",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "K", Weight: 13},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isFlush(), test)
	}
}

func TestHand_isStraight(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Straight",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "2", Weight: 2},
				{Suit: "D", Name: "3", Weight: 3},
				{Suit: "S", Name: "4", Weight: 4},
				{Suit: "H", Name: "5", Weight: 5},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Straight",
				HandWeight:        28,
				CombinationWeight: straightCombinationWeight,
			},
		},
		{
			name: "Ace-Low Straight",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "2", Weight: 2},
				{Suit: "D", Name: "3", Weight: 3},
				{Suit: "S", Name: "4", Weight: 4},
				{Suit: "H", Name: "5", Weight: 5},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Straight",
				HandWeight:        28,
				CombinationWeight: straightCombinationWeight,
			},
		},
		{
			name: "No Straight",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "K", Weight: 13},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isStraight(), test)
	}
}

func TestHand_isThreeOfAKind(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Three of a Kind",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "A", Weight: 14},
				{Suit: "D", Name: "A", Weight: 14},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Three of a kind",
				HandWeight:        54,
				CombinationWeight: threeOfAKindCombinationWeight,
			},
		},
		{
			name: "No Three of a Kind",
			cards: []Card{
				{Suit: "H", Name: "A", Weight: 14},
				{Suit: "C", Name: "K", Weight: 13},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "5", Weight: 5},
				{Suit: "H", Name: "7", Weight: 7},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isThreeOfAKind(), test)
	}
}

func TestHand_isTwoPair(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Two Pair",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "9", Weight: 9},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "J", Weight: 11},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Two pair",
				HandWeight:        53,
				CombinationWeight: twoPairCombinationWeight,
			},
		},
		{
			name: "No Two Pair (Single Pair)",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "9", Weight: 9},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: nil,
		},
		{
			name: "No Two Pair (No Pairs)",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isTwoPair(), test)
	}
}

func TestHand_isPair(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Pair",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "9", Weight: 9},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Pair",
				HandWeight:        54,
				CombinationWeight: pairCombinationWeight,
			},
		},
		{
			name: "No Pair",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isPair(), test)
	}
}

func TestHand_isHighCard(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "High Card",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "D", Name: "Q", Weight: 12},
				{Suit: "S", Name: "K", Weight: 13},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "High card",
				HandWeight:        55,
				CombinationWeight: highCardCombinationWeight,
			},
		},
		{
			name: "High Card (All Same Suit)",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "High card",
				HandWeight:        55,
				CombinationWeight: highCardCombinationWeight,
			},
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.isHighCard(), test)
	}
}

func TestHand_DefineCombination(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected *HandResult
	}{
		{
			name: "Royal Flush",
			cards: []Card{
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
				{Suit: "H", Name: "Q", Weight: 12},
				{Suit: "H", Name: "K", Weight: 13},
				{Suit: "H", Name: "A", Weight: 14},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Royal Flush",
				HandWeight:        60,
				CombinationWeight: royalFlushCombinationWeight,
			},
		},
		{
			name: "Straight Flush",
			cards: []Card{
				{Suit: "H", Name: "7", Weight: 7},
				{Suit: "H", Name: "8", Weight: 8},
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "H", Name: "T", Weight: 10},
				{Suit: "H", Name: "J", Weight: 11},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Straight Flush",
				HandWeight:        45,
				CombinationWeight: straightFlushCombinationWeight,
			},
		},
		{
			name: "Four of a Kind",
			cards: []Card{
				{Suit: "H", Name: "9", Weight: 9},
				{Suit: "C", Name: "9", Weight: 9},
				{Suit: "D", Name: "9", Weight: 9},
				{Suit: "S", Name: "9", Weight: 9},
				{Suit: "H", Name: "J", Weight: 11},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "Four of a kind",
				HandWeight:        47,
				CombinationWeight: fourOfAKindCombinationWeight,
			},
		},
		{
			name: "High Card",
			cards: []Card{
				{Suit: "H", Name: "2", Weight: 2},
				{Suit: "C", Name: "3", Weight: 3},
				{Suit: "H", Name: "5", Weight: 5},
				{Suit: "D", Name: "J", Weight: 11},
				{Suit: "S", Name: "A", Weight: 14},
			},
			expected: &HandResult{
				HandName:          "Test Hand",
				CombinationName:   "High card",
				HandWeight:        35,
				CombinationWeight: highCardCombinationWeight,
			},
		},
	}

	for _, test := range tests {
		hand := &Hand{
			Name:  "Test Hand",
			Cards: test.cards,
		}
		testHandResult(t, hand.DefineCombination(), test)
	}
}

func testHandResult(
	t *testing.T,
	result *HandResult,
	test struct {
		name     string
		cards    []Card
		expected *HandResult
	}) {
	t.Run(test.name, func(t *testing.T) {
		if (result == nil && test.expected != nil) || (result != nil && test.expected == nil) {
			t.Errorf("Expected result %v, but got %v", test.expected, result)
			return
		}
		if result != nil && test.expected != nil {
			if result.CombinationName != test.expected.CombinationName ||
				result.HandWeight != test.expected.HandWeight ||
				result.CombinationWeight != test.expected.CombinationWeight {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		}
	})
}
