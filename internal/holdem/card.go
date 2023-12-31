package holdem

type CardSuit string
type CardName string
type CardWeight int32

var (
	// Using array for cards representation because we already know cards count. So we can optimize some memory consumption.
	cardsList = [13]CardName{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	// Map are used for convenient weights calculations.
	cardWeights = map[CardName]CardWeight{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
)

type Card struct {
	Name   CardName
	Suit   CardSuit
	Weight CardWeight
}

func (c *Card) ResolveWeight() CardWeight {
	return cardWeights[c.Name]
}

func ResolveWeight(cardName CardName) CardWeight {
	return cardWeights[cardName]
}
