package tetris

type RockGenerator struct {
	n       int
	created int
}

var creationSequence = []Rock{rock1, rock2, rock3, rock4, rock5}

func NewRockGenerator() *RockGenerator {
	return &RockGenerator{
		n: len(creationSequence),
	}
}

func (r *RockGenerator) Generate() *Rock {
	idx := r.created % r.n
	r.created++

	return NewRock(creationSequence[idx])
}

type JetGenerator struct {
	idx     int
	max     int
	pattern []byte
}

func NetJetGenerator(pattern []byte) *JetGenerator {
	return &JetGenerator{
		pattern: pattern,
		max:     len(pattern),
	}
}

var idx = 0

func (r *JetGenerator) Generate() byte {
	jet := r.pattern[r.idx]

	r.idx++
	if r.idx == r.max {
		r.idx = 0
	}
	idx = r.idx

	return jet
}
