package holdem

import (
	"sort"
)

const (
	royalFlushCombinationWeight    = 10
	straightFlushCombinationWeight = 9
	fourOfAKindCombinationWeight   = 8
	fullHouseCombinationWeight     = 7
	flushCombinationWeight         = 6
	straightCombinationWeight      = 5
	threeOfAKindCombinationWeight  = 4
	twoPairCombinationWeight       = 3
	pairCombinationWeight          = 2
	highCardCombinationWeight      = 1
)

type Hands map[string][]string

type HandResult struct {
	HandName          string `json:"handName"`
	CombinationName   string `json:"combinationName"`
	HandWeight        int32  `json:"handWeight"`
	CombinationWeight int32  `json:"combinationWeight"`
}

type Hand struct {
	Name  string
	Cards []Card
}

func (h *Hand) DefineCombination() *HandResult {
	if royalFlushResult := h.isRoyalFlush(); royalFlushResult != nil {
		return royalFlushResult
	}

	if straightFlushResult := h.isStraightFlush(); straightFlushResult != nil {
		return straightFlushResult
	}

	if fourOfAKindResult := h.isFourOfAKind(); fourOfAKindResult != nil {
		return fourOfAKindResult
	}

	if fullHouseResult := h.isFullHouse(); fullHouseResult != nil {
		return fullHouseResult
	}

	if flushResult := h.isFlush(); flushResult != nil {
		return flushResult
	}

	if straightResult := h.isStraight(); straightResult != nil {
		return straightResult
	}

	if threeOfAKindResult := h.isThreeOfAKind(); threeOfAKindResult != nil {
		return threeOfAKindResult
	}

	if twoPairResult := h.isTwoPair(); twoPairResult != nil {
		return twoPairResult
	}

	if pairResult := h.isPair(); pairResult != nil {
		return pairResult
	}

	return h.isHighCard()
}

func (h *Hand) calculateHandWeight() CardWeight {
	var totalWeight CardWeight
	for _, card := range h.Cards {
		totalWeight += card.ResolveWeight()
	}
	return totalWeight
}

func (h *Hand) isRoyalFlush() *HandResult {
	firstSuit := h.Cards[0].Suit
	allCardsSame := true
	for _, card := range h.Cards {
		if card.Suit != firstSuit {
			allCardsSame = false
			break
		}
	}

	if !allCardsSame {
		return nil
	}

	royalFlushCombination := []CardName{"T", "J", "Q", "K", "A"}
	for i, cardName := range royalFlushCombination {
		if i >= len(h.Cards) || h.Cards[i].Name != cardName {
			return nil
		}
	}

	return &HandResult{
		HandName:          h.Name,
		CombinationName:   "Royal Flush",
		HandWeight:        int32(h.calculateHandWeight()),
		CombinationWeight: royalFlushCombinationWeight,
	}
}

func (h *Hand) isStraightFlush() *HandResult {
	suitCount := make(map[CardSuit]int)

	for _, card := range h.Cards {
		suitCount[card.Suit]++
	}

	var potentialSuit CardSuit
	for suit, count := range suitCount {
		if count >= 5 {
			potentialSuit = suit
			break
		}
	}

	if potentialSuit == "" {
		return nil // No suit occurs at least five times, no Straight Flush possible
	}

	potentialSuitCards := make([]Card, 0, 7)
	for _, card := range h.Cards {
		if card.Suit == potentialSuit {
			potentialSuitCards = append(potentialSuitCards, card)
		}
	}
	sort.Slice(potentialSuitCards, func(i, j int) bool {
		return potentialSuitCards[i].Weight < potentialSuitCards[j].Weight
	})

	// Check if there is a sequence of five consecutive cards
	consecutiveCount := 1
	for i := 1; i < len(potentialSuitCards); i++ {
		if potentialSuitCards[i].Weight-potentialSuitCards[i-1].Weight == 1 {
			consecutiveCount++
			if consecutiveCount == 5 {
				return &HandResult{
					HandName:          h.Name,
					CombinationName:   "Straight Flush",
					HandWeight:        int32(h.calculateHandWeight()),
					CombinationWeight: straightFlushCombinationWeight,
				}
			}
		} else {
			consecutiveCount = 1
		}
	}

	return nil
}

func (h *Hand) isFullHouse() *HandResult {
	rankCount := make(map[CardName]int)
	for _, card := range h.Cards {
		rankCount[card.Name]++
	}

	var hasThreeOfAKind, hasPair bool

	for _, count := range rankCount {
		switch count {
		case 3:
			hasThreeOfAKind = true
		case 2:
			hasPair = true
		}
	}

	if hasThreeOfAKind && hasPair {
		return &HandResult{
			HandName:          h.Name,
			CombinationName:   "Full House",
			HandWeight:        int32(h.calculateHandWeight()),
			CombinationWeight: fullHouseCombinationWeight,
		}
	}

	return nil
}

func (h *Hand) isFourOfAKind() *HandResult {
	rankCount := make(map[CardName]int)
	for _, card := range h.Cards {
		rankCount[card.Name]++
	}
	for _, count := range rankCount {
		if count == 4 {
			return &HandResult{
				HandName:          h.Name,
				CombinationName:   "Four of a kind",
				HandWeight:        int32(h.calculateHandWeight()),
				CombinationWeight: fourOfAKindCombinationWeight,
			}
		}
	}
	return nil
}

func (h *Hand) isFlush() *HandResult {
	suitCount := make(map[CardSuit]int)
	maxRank := CardName("2")

	for _, card := range h.Cards {
		suitCount[card.Suit]++
		if card.Suit == h.Cards[0].Suit && card.Name > maxRank {
			maxRank = card.Name
		}
	}

	for _, count := range suitCount {
		if count >= 5 {
			flushSuit := h.Cards[0].Suit
			var flushCards []Card
			for _, card := range h.Cards {
				if card.Suit == flushSuit {
					flushCards = append(flushCards, card)
				}
			}
			sort.Slice(flushCards, func(i, j int) bool {
				return flushCards[i].Name > flushCards[j].Name
			})
			return &HandResult{
				HandName:          h.Name,
				CombinationName:   "Flush",
				HandWeight:        int32(h.calculateHandWeight()),
				CombinationWeight: flushCombinationWeight,
			}
		}
	}

	return nil
}

func (h *Hand) isStraight() *HandResult {
	rankCount := make(map[CardName]int)
	uniqueRanks := make(map[CardName]bool)

	for _, card := range h.Cards {
		rankCount[card.Name]++
		uniqueRanks[card.Name] = true
	}

	var consecutiveCount int

	// Check for Ace as low card (A-2-3-4-5 straight)
	if uniqueRanks["A"] && uniqueRanks["2"] && uniqueRanks["3"] && uniqueRanks["4"] && uniqueRanks["5"] {
		consecutiveCount++
	}

	for _, cardName := range cardsList {
		if rankCount[cardName] > 0 {
			consecutiveCount++
			if consecutiveCount == 5 {
				return &HandResult{
					HandName:          h.Name,
					CombinationName:   "Straight",
					HandWeight:        int32(h.calculateHandWeight()),
					CombinationWeight: straightCombinationWeight,
				}
			}
		} else {
			consecutiveCount = 0
		}
	}

	return nil
}

func (h *Hand) isThreeOfAKind() *HandResult {
	rankCount := make(map[CardName]int)

	for _, card := range h.Cards {
		rankCount[card.Name]++
	}

	for _, count := range rankCount {
		if count == 3 {
			return &HandResult{
				HandName:          h.Name,
				CombinationName:   "Three of a kind",
				HandWeight:        int32(h.calculateHandWeight()),
				CombinationWeight: threeOfAKindCombinationWeight,
			}
		}
	}

	return nil
}

func (h *Hand) isTwoPair() *HandResult {
	rankCount := make(map[CardName]int)
	pairCount := 0

	for _, card := range h.Cards {
		rankCount[card.Name]++
		if rankCount[card.Name] == 2 {
			pairCount++
		}
	}

	if pairCount >= 2 {
		return &HandResult{
			HandName:          h.Name,
			CombinationName:   "Two pair",
			HandWeight:        int32(h.calculateHandWeight()),
			CombinationWeight: twoPairCombinationWeight,
		}
	}
	return nil
}

func (h *Hand) isPair() *HandResult {
	rankCount := make(map[CardName]int)

	for _, card := range h.Cards {
		rankCount[card.Name]++
	}

	for _, count := range rankCount {
		if count == 2 {
			return &HandResult{
				HandName:          h.Name,
				CombinationName:   "Pair",
				HandWeight:        int32(h.calculateHandWeight()),
				CombinationWeight: pairCombinationWeight,
			}
		}
	}

	return nil
}

func (h *Hand) isHighCard() *HandResult {
	rankCount := make(map[CardName]int)

	for _, card := range h.Cards {
		rankCount[card.Name]++
	}

	if len(rankCount) == 5 {
		return &HandResult{
			HandName:          h.Name,
			CombinationName:   "High card",
			HandWeight:        int32(h.calculateHandWeight()),
			CombinationWeight: highCardCombinationWeight,
		}
	}

	return nil
}
