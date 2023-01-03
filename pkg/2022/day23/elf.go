package day23

var around = []string{N, S, E, W, NW, NE, SW, SE}

var dirGroups = [][]string{
	{N, NE, NW},
	{S, SE, SW},
	{W, NW, SW},
	{E, NE, SE},
}

var proposalByGroup = map[int]string{
	0: N,
	1: S,
	2: W,
	3: E,
}

type Elf struct {
	Position      Coordinate
	groupIdx      int
	proposedGroup int
	proposedPos   Coordinate
}

func (elf *Elf) Propose(grove *Grove) (bool, Coordinate) {
	var freeGroups []int
	idx := elf.groupIdx
	for range dirGroups {
		freeGroup := true
		for _, p := range dirGroups[idx] {
			check := elf.Position.Add(increment[p])
			if grove.Occupied(check) {
				freeGroup = false
				break
			}
		}
		if freeGroup {
			freeGroups = append(freeGroups, idx)
		}

		idx = (idx + 1) & 0x3
	}

	if len(freeGroups) == 4 {
		return false, Coordinate{}
	}

	if len(freeGroups) == 0 {
		return false, Coordinate{}
	}

	idx = freeGroups[0]
	elf.proposedPos = elf.Position.Add(increment[proposalByGroup[idx]])

	return true, elf.proposedPos
}

func (elf *Elf) AcceptProposal() {
	elf.Position = elf.proposedPos
}

func (elf *Elf) EndRound() {
	elf.groupIdx = (elf.groupIdx + 1) & 0x3
}
