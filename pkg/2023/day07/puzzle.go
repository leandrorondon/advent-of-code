package day07

import (
	"slices"
	"sort"
	"strings"
)

type HandType int
type counterFunc func(s string, values map[rune]int) map[int]int

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

func (h Hand) Against(other Hand, values map[rune]int) int {
	if h.Type > other.Type {
		return 1
	}
	if h.Type < other.Type {
		return -1
	}
	for i := 0; i < 5; i++ {
		value1 := values[rune(h.Hand[i])]
		value2 := values[rune(other.Hand[i])]
		if value1 > value2 {
			return 1
		} else if value1 < value2 {
			return -1
		}
	}
	return 0
}

var CardValuesP1 = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

var CardValuesP2 = map[rune]int{
	'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, 'T': 10, 'Q': 12, 'K': 13, 'A': 14,
}

func GetCardsAndCountsP1(s string, values map[rune]int) map[int]int {
	counts := make(map[int]int)
	var cards []int
	for _, card := range s {
		value := values[card]
		counts[value]++
		cards = append(cards, value)
	}
	return counts
}

func GetCardsAndCountsP2(s string, values map[rune]int) map[int]int {
	counts := make(map[int]int)
	var cards []int
	highest := 0
	valueHighest := 0
	jValue := CardValuesP2['J']
	for _, card := range s {
		value := values[card]
		counts[value]++

		if value != jValue && (counts[value] > highest || (counts[value] == highest && value > valueHighest)) {
			highest = counts[value]
			valueHighest = value
		}

		cards = append(cards, value)
	}

	// joker
	if counts[jValue] > 0 {
		counts[valueHighest] += counts[jValue]
		delete(counts, jValue)
	}

	return counts
}

func getHandType(counts map[int]int, rev map[int]rune) (HandType, string) {
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
		orderedCards = strings.Repeat(string(rev[fives[0]]), 5)
	case len(fours) == 1:
		handType = FourOfAKind
		orderedCards = strings.Repeat(string(rev[fours[0]]), 4) +
			string(rev[ones[0]])
	case len(threes) == 1 && len(pairs) == 1:
		handType = FullHouse
		orderedCards = strings.Repeat(string(rev[threes[0]]), 3) +
			strings.Repeat(string(rev[pairs[0]]), 2)
	case len(threes) == 1:
		handType = ThreeOfAKind
		orderedCards = strings.Repeat(string(rev[threes[0]]), 3) +
			string(rev[ones[0]]) +
			string(rev[ones[1]])

	case len(pairs) == 2:
		handType = TwoPair
		orderedCards = strings.Repeat(string(rev[pairs[0]]), 2) +
			strings.Repeat(string(rev[pairs[1]]), 2) +
			string(rev[ones[0]])
	case len(pairs) == 1:
		handType = OnePair
		orderedCards = strings.Repeat(string(rev[pairs[0]]), 2) +
			string(rev[ones[0]]) +
			string(rev[ones[1]]) +
			string(rev[ones[2]])
	default:
		handType = HighCard
		orderedCards = string(rev[ones[0]]) +
			string(rev[ones[1]]) +
			string(rev[ones[2]]) +
			string(rev[ones[3]]) +
			string(rev[ones[4]])
	}

	return handType, orderedCards
}
func NewHand(s string, values map[rune]int, rev map[int]rune, counter counterFunc) Hand {
	counts := counter(s, values)
	handType, ordered := getHandType(counts, rev)
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

type Plays struct {
	plays  []Play
	values map[rune]int
}

func (p Plays) Len() int {
	return len(p.plays)
}

func (p Plays) Less(i, j int) bool {
	return p.plays[i].Hand.Against(p.plays[j].Hand, p.values) <= 0
}

func (p Plays) Swap(i, j int) {
	p.plays[i], p.plays[j] = p.plays[j], p.plays[i]
}

func (p Plays) TotalWins() int {
	sort.Sort(p)
	total := 0
	for i := range p.plays {
		total += (i + 1) * p.plays[i].Bid
	}
	return total
}
