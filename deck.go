package cards

import (
	"errors"
	"math/rand"
)

type Deck []Card

var ErrHandSizeTooLarge = errors.New("hand size cannot be larger than the deck size")
var ErrHandSizeTooSmall = errors.New("hand size must be positive")

func NewDeck(includeJokers bool) Deck {
	var deck Deck
	for suit := Spades; suit <= Clubs; suit++ {
		for rank := Ace; rank <= King; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	if includeJokers {
		deck = append(deck, Card{Suit: JokerSuit, Rank: Joker1Rank}, Card{Suit: JokerSuit, Rank: Joker2Rank})
	}

	return deck
}

func Shuffle(d Deck, r *rand.Rand) Deck {
	shuffledDeck := make(Deck, len(d))
	copy(shuffledDeck, d)

	r.Shuffle(len(shuffledDeck), func(i, j int) {
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	})

	return shuffledDeck
}

func Deal(d Deck, handSize int) (hand Deck, remainingDeck Deck, err error) {
	if handSize > len(d) {
		return nil, d, ErrHandSizeTooLarge
	}

	if handSize <= 0 {
		return nil, d, ErrHandSizeTooSmall
	}

	hand = d[:handSize]
	remainingDeck = d[handSize:]

	return hand, remainingDeck, nil
}
