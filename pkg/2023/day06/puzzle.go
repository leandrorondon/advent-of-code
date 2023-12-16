package day06

type Race struct {
	Time     int
	Distance int
}

func (r Race) WaysToWin() (wins int) {
	for i := 1; i < r.Time; i++ {
		if r.DistanceWithPress(i) > r.Distance {
			wins++
		}
	}
	return wins
}

func (r Race) DistanceWithPress(press int) int {
	return (r.Time - press) * press
}

type Races []Race

func (rs Races) MarginOfError() int {
	margin := 1
	for _, r := range rs {
		margin *= r.WaysToWin()
	}
	return margin
}
