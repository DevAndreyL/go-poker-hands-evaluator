package holdem

type EvaluateResult struct {
	Result map[string]*HandResult `json:"result"`
}

func EvaluateAndCompareHands(hands Hands) (*EvaluateResult, error) {
	handsWithCards := createCardsForHands(hands)

	handCombinations := map[string]*HandResult{}

	for _, hand := range handsWithCards {
		handCombinations[hand.Name] = hand.DefineCombination()
	}

	return &EvaluateResult{
		Result: handCombinations,
	}, nil
}

func createCardsForHands(hands Hands) []Hand {
	var result []Hand
	for handName, handCards := range hands {
		var cardsInHand []Card
		for _, cardString := range handCards {
			name := CardName(cardString[0])
			cardsInHand = append(cardsInHand, Card{
				Name:   name,
				Suit:   CardSuit(cardString[1]),
				Weight: ResolveWeight(name),
			})
		}
		result = append(result, Hand{
			Name:  handName,
			Cards: cardsInHand,
		})
	}
	return result
}
