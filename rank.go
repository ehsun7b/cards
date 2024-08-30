package cards

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Joker1Rank
	Joker2Rank
)

func (r Rank) String() string {
	switch r {
	case Joker1Rank:
		return "Joker 1"
	case Joker2Rank:
		return "Joker 2"
	default:
		return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r-1]
	}
}
