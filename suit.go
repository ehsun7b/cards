package cards

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
	JokerSuit
)

func (s Suit) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs", "Joker"}[s]
}
