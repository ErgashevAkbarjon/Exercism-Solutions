package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":   11,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
	}

	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch true {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum == 21 && dealerCardValue < 10:
		return "W"
	case sum == 21 && dealerCardValue >= 10:
		return "S"
	case sum >= 17 && sum <= 21:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardValue < 7:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardValue >= 7:
		return "H"
	case sum <= 11:
		return "H"
	default:
		return "H" // Default case, should not be reached
	}
}
