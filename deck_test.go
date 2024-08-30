package cards

import (
	"math/rand"
	"testing"
)

func TestNewDeckWithDistinguishableJokers(t *testing.T) {
	deck := NewDeck(true)

	// Test the length of the deck
	expectedLength := 54
	if len(deck) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(deck))
	}

	// Test the first card
	expectedFirstCard := Card{Suit: Spades, Rank: Ace}
	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, deck[0])
	}

	// Test the last two cards (distinguishable jokers)
	expectedJoker1 := Card{Suit: JokerSuit, Rank: Joker1Rank}
	expectedJoker2 := Card{Suit: JokerSuit, Rank: Joker2Rank}
	if deck[len(deck)-2] != expectedJoker1 || deck[len(deck)-1] != expectedJoker2 {
		t.Errorf("Expected the last two cards to be %v and %v, but got %v and %v", expectedJoker1, expectedJoker2, deck[len(deck)-2], deck[len(deck)-1])
	}
}

func TestNewDeckWithoutJokers(t *testing.T) {
	deck := NewDeck(false)

	// Test the length of the deck
	expectedLength := 52
	if len(deck) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(deck))
	}

	// Test the first card
	expectedFirstCard := Card{Suit: Spades, Rank: Ace}
	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, deck[0])
	}

	// Test the last card
	expectedLastCard := Card{Suit: Clubs, Rank: King}
	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, deck[len(deck)-1])
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck(true)

	// Create a random source with a fixed seed for reproducibility
	r := rand.New(rand.NewSource(42))

	shuffledDeck := Shuffle(deck, r)

	// Test that the shuffled deck has the same length as the original
	if len(shuffledDeck) != len(deck) {
		t.Errorf("Expected shuffled deck length to be %d, but got %d", len(deck), len(shuffledDeck))
	}

	// Test that the shuffled deck is not in the same order as the original deck
	isSameOrder := true
	for i, card := range deck {
		if shuffledDeck[i] != card {
			isSameOrder = false
			break
		}
	}
	if isSameOrder {
		t.Errorf("Expected the shuffled deck to be in a different order, but it is the same as the original")
	}

	// Test that all original cards are still present in the shuffled deck
	cardMap := make(map[Card]bool)
	for _, card := range deck {
		cardMap[card] = true
	}
	for _, card := range shuffledDeck {
		if !cardMap[card] {
			t.Errorf("Shuffled deck is missing card: %v", card)
		}
	}
}

func TestShuffleWithDifferentSeeds(t *testing.T) {
	deck := NewDeck(true)

	// Shuffle with one seed
	r1 := rand.New(rand.NewSource(42))
	shuffledDeck1 := Shuffle(deck, r1)

	// Shuffle with a different seed
	r2 := rand.New(rand.NewSource(43))
	shuffledDeck2 := Shuffle(deck, r2)

	// Test that shuffling with different seeds produces different orders
	if isSameOrder(shuffledDeck1, shuffledDeck2) {
		t.Errorf("Expected shuffled decks with different seeds to be in different orders, but they are the same")
	}
}

// Helper function to check if two decks are in the same order
func isSameOrder(deck1, deck2 Deck) bool {
	if len(deck1) != len(deck2) {
		return false
	}
	for i := range deck1 {
		if deck1[i] != deck2[i] {
			return false
		}
	}
	return true
}

func TestDealNegativeHandSize(t *testing.T) {
	deck := NewDeck(false) // 52 cards without jokers

	handSize := -1
	_, _, err := Deal(deck, handSize)
	if err == nil {
		t.Errorf("Expected an error for negative hand size, but got no error")
	}
}

func TestDealZeroHandSize(t *testing.T) {
	deck := NewDeck(false) // 52 cards without jokers

	handSize := 0
	_, _, err := Deal(deck, handSize)
	if err == nil {
		t.Errorf("Expected an error for zero hand size, but got no error")
	}
}

func TestDealHandSizeTooLarge(t *testing.T) {
	deck := NewDeck(false) // 52 cards without jokers

	handSize := 60
	_, _, err := Deal(deck, handSize)
	if err == nil {
		t.Errorf("Expected an error when hand size is larger than deck size, but got no error")
	} else if err != ErrHandSizeTooLarge {
		t.Errorf("Expected error '%v', but got '%v'", ErrHandSizeTooLarge, err)
	}
}

func TestDealHandSizeTooSmall(t *testing.T) {
	deck := NewDeck(false) // 52 cards without jokers

	handSizes := []int{0, -1, -10}

	for _, handSize := range handSizes {
		_, _, err := Deal(deck, handSize)
		if err == nil {
			t.Errorf("Expected an error for hand size %d, but got no error", handSize)
		} else if err != ErrHandSizeTooSmall {
			t.Errorf("Expected error '%v', but got '%v'", ErrHandSizeTooSmall, err)
		}
	}
}

func TestRemoveCard(t *testing.T) {
	deck := NewDeck(true) // A deck with jokers

	cardsToRemove := Deck{
		{Suit: JokerSuit, Rank: Joker1Rank},
		{Suit: JokerSuit, Rank: Joker2Rank},
	}

	remainingDeck := RemoveCard(deck, cardsToRemove)

	// Verify that the remaining deck does not contain the removed cards
	for _, card := range cardsToRemove {
		for _, remainingCard := range remainingDeck {
			if card == remainingCard {
				t.Errorf("Expected card %v to be removed, but it was found in the remaining deck", card)
			}
		}
	}

	// Verify the length of the remaining deck
	expectedLength := len(deck) - len(cardsToRemove)
	if len(remainingDeck) != expectedLength {
		t.Errorf("Expected remaining deck length to be %d, but got %d", expectedLength, len(remainingDeck))
	}
}
