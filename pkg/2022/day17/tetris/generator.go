package tetris

type RockGenerator struct {
	idx int
	max int
}

var creationSequence = []Rock{rock1, rock2, rock3, rock4, rock5}

func NewRockGenerator() *RockGenerator {
	return &RockGenerator{
		max: len(creationSequence),
	}
}

func (g *RockGenerator) Generate() *Rock {
	rock := creationSequence[g.idx]

	g.idx++
	if g.idx == g.max {
		g.idx = 0
	}
	idx = g.idx

	return NewRock(rock)
}

func (g *RockGenerator) Next() int {
	return g.idx
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

func (g *JetGenerator) Generate() byte {
	jet := g.pattern[g.idx]

	g.idx++
	if g.idx == g.max {
		g.idx = 0
	}
	idx = g.idx

	return jet
}

func (g *JetGenerator) Next() int {
	return g.idx
}
