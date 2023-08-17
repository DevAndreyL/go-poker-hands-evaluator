package holdem

type EvaluateResult struct {
	//TODO
}

func EvaluateAndCompareHands(hands Hands) (*EvaluateResult, error) {
	handsWithCards := createCardsForHands(hands)
	for _, hand := range handsWithCards {
		hand.DefineCombination()
	}
	return &EvaluateResult{}, nil
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
