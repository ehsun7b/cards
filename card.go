package cards

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return c.Rank.String() + " of " + c.Suit.String()
}
