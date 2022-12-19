package tetris

func Simulate(chamber *Chamber, n int) {
	for i := 0; i < n; i++ {
		chamber.SpawnRock()
		for !chamber.CurrentRockResting() {
			chamber.Jet()
			chamber.Tick()
		}
	}
}

type DeepPattern struct {
	Rocks    int
	Height   int
	Depths   []int
	NextRock int
	NextJet  int
}

func (p *DeepPattern) SamePattern(other *DeepPattern) bool {
	if p == nil || other == nil {
		return false
	}

	if p.NextJet != other.NextJet {
		return false
	}

	if p.NextRock != other.NextRock {
		return false
	}

	if len(p.Depths) != len(other.Depths) {
		return false
	}

	for i := range p.Depths {
		if p.Depths[i] != other.Depths[i] {
			return false
		}
	}

	return true
}

func FindSimilarPatterns(chamber *Chamber) (*DeepPattern, *DeepPattern) {
	var patterns []*DeepPattern
	rocks := 0

	for {
		p := buildPattern(chamber, rocks)
		if found := search(p, patterns); found != nil {
			return found, p
		}

		patterns = append(patterns, p)

		chamber.SpawnRock()
		rocks++
		for !chamber.CurrentRockResting() {
			chamber.Jet()
			chamber.Tick()
		}
	}

	return nil, nil
}

func search(pattern *DeepPattern, patterns []*DeepPattern) *DeepPattern {
	for _, p := range patterns {
		if p.SamePattern(pattern) {
			return p
		}
	}
	return nil
}

func buildPattern(chamber *Chamber, rocks int) *DeepPattern {
	found := 0 // bit mask for each column
	depths := make([]int, chamber.Width)
	for y := chamber.Height - 1; y >= 0; y-- {
		for i := 0; i < chamber.Width; i++ {
			bit := 1 << i
			if bit&found == 0 && chamber.chamber[y]&bit > 0 {
				found |= bit
				depths[i] = chamber.Height - y
			}
		}
	}

	return &DeepPattern{
		Rocks:    rocks,
		Height:   chamber.Height,
		Depths:   depths,
		NextRock: chamber.NextRockIndex(),
		NextJet:  chamber.NextJetIndex(),
	}
}
