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

func RemoveCards(d Deck, cardsToRemove Deck) Deck {
	toRemoveMap := make(map[Card]bool)
	for _, card := range cardsToRemove {
		toRemoveMap[card] = true
	}

	var remainingDeck Deck
	for _, card := range d {
		if !toRemoveMap[card] {
			remainingDeck = append(remainingDeck, card)
		}
	}

	return remainingDeck
}

// with receiver versions

func (d *Deck) Shuffle(r *rand.Rand) {
	shuffled := Shuffle(*d, r)
	*d = shuffled
}

func (d *Deck) Deal(handSize int) (hand Deck, err error) {
	hand, remainingDeck, err := Deal(*d, handSize)
	if err != nil {
		return nil, err
	}

	*d = remainingDeck

	return hand, nil
}

func (d *Deck) RemoveCards(cardsToRemove Deck) {
	remainingDeck := RemoveCards(*d, cardsToRemove)

	*d = remainingDeck
}
