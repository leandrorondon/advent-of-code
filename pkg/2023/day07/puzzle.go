package day07

import (
	"slices"
	"sort"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Hand        string
	Type        HandType
	orderedHand string
}

func (h Hand) Against(other Hand) int {
	if h.Type > other.Type {
		return 1
	}
	if h.Type < other.Type {
		return -1
	}
	for i := 0; i < 5; i++ {
		value1 := cardValues[rune(h.Hand[i])]
		value2 := cardValues[rune(other.Hand[i])]
		if value1 > value2 {
			return 1
		} else if value1 < value2 {
			return -1
		}
	}
	return 0
}

var cardValues = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

var reversedCardValues = map[int]rune{
	2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8',
	9: '9', 10: 'T', 11: 'J', 12: 'Q', 13: 'K', 14: 'A',
}

func getCardsAndCounts(s string) map[int]int {
	counts := make(map[int]int)
	var cards []int
	for _, card := range s {
		value := cardValues[card]

		counts[value]++
		cards = append(cards, value)
	}
	return counts
}

func getHandType(counts map[int]int) (HandType, string) {
	var handType HandType
	var ones, pairs, threes, fours, fives []int

	for card, count := range counts {
		switch count {
		case 1:
			ones = append(ones, card)
		case 2:
			pairs = append(pairs, card)
		case 3:
			threes = append(threes, card)
		case 4:
			fours = append(fours, card)
		case 5:
			fives = append(fives, card)
		}
	}

	slices.Sort(fives)
	slices.Sort(fours)
	slices.Sort(threes)
	slices.Sort(pairs)
	slices.Sort(ones)
	slices.Reverse(pairs)
	slices.Reverse(ones)

	var orderedCards string

	switch {
	case len(fives) == 1:
		handType = FiveOfAKind
		orderedCards = strings.Repeat(string(reversedCardValues[fives[0]]), 5)
	case len(fours) == 1:
		handType = FourOfAKind
		orderedCards = strings.Repeat(string(reversedCardValues[fours[0]]), 4) +
			string(reversedCardValues[ones[0]])
	case len(threes) == 1 && len(pairs) == 1:
		handType = FullHouse
		orderedCards = strings.Repeat(string(reversedCardValues[threes[0]]), 3) +
			strings.Repeat(string(reversedCardValues[pairs[0]]), 2)
	case len(threes) == 1:
		handType = ThreeOfAKind
		orderedCards = strings.Repeat(string(reversedCardValues[threes[0]]), 3) +
			string(reversedCardValues[ones[0]]) +
			string(reversedCardValues[ones[1]])

	case len(pairs) == 2:
		handType = TwoPair
		orderedCards = strings.Repeat(string(reversedCardValues[pairs[0]]), 2) +
			strings.Repeat(string(reversedCardValues[pairs[1]]), 2) +
			string(reversedCardValues[ones[0]])
	case len(pairs) == 1:
		handType = OnePair
		orderedCards = strings.Repeat(string(reversedCardValues[pairs[0]]), 2) +
			string(reversedCardValues[ones[0]]) +
			string(reversedCardValues[ones[1]]) +
			string(reversedCardValues[ones[2]])
	default:
		handType = HighCard
		orderedCards = string(reversedCardValues[ones[0]]) +
			string(reversedCardValues[ones[1]]) +
			string(reversedCardValues[ones[2]]) +
			string(reversedCardValues[ones[3]]) +
			string(reversedCardValues[ones[4]])
	}

	return handType, orderedCards
}
func NewHand(s string) Hand {
	counts := getCardsAndCounts(s)
	handType, ordered := getHandType(counts)
	sort.Sort(Plays{})

	return Hand{
		Hand:        s,
		Type:        handType,
		orderedHand: ordered,
	}
}

type Play struct {
	Hand Hand
	Bid  int
}

type Plays []Play

func (p Plays) Len() int {
	return len(p)
}

func (p Plays) Less(i, j int) bool {
	return p[i].Hand.Against(p[j].Hand) <= 0
}

func (p Plays) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Plays) TotalWins() int {
	sort.Sort(p)
	total := 0
	for i := range p {
		total += (i + 1) * p[i].Bid
	}
	return total
}
