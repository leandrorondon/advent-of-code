package decoder

import "sort"

type Signal []Packet

func (s Signal) PairsInRightOrder() int {
	sum := 0
	pair := 0
	for i := 0; i < len(s)-1; i = i + 2 {
		pair++
		p1 := s[i]
		p2 := s[i+1]
		cmp := p1.Compare(p2)
		if cmp < 0 {
			sum += pair
		}
	}
	return sum
}

func (s Signal) DecoderKey(divider1, divider2 Packet) int {
	s = append(s, divider1)
	s = append(s, divider2)

	sort.Slice(s, func(i, j int) bool {
		return s[i].Compare(s[j]) < 0
	})

	decoderKey := 1
	for i := 0; i < len(s); i++ {
		if s[i].Compare(divider1) == 0 ||
			s[i].Compare(divider2) == 0 {
			decoderKey *= i + 1
		}
	}

	return decoderKey
}
