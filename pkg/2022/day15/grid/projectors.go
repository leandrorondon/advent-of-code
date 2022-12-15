package grid

import mymath "github.com/leandrorondon/advent-of-code/pkg/math"

type SimpleProjector struct{}

func NewSimpleProjector() SimpleProjector {
	return SimpleProjector{}
}

func (p SimpleProjector) ProjectInRow(sensor Sensor, row int) *Interval {
	rowDist := mymath.Abs(row - sensor.Position.Y)

	if rowDist > sensor.Distance {
		// Out of range
		return nil
	}

	n := sensor.Distance - rowDist
	from := sensor.Position.X - n
	to := sensor.Position.X + n

	return &Interval{From: from, To: to}
}

type LimitedProjector struct {
	max int
}

func NewLimitedProjector(max int) LimitedProjector {
	return LimitedProjector{
		max: max,
	}
}

func (p LimitedProjector) ProjectInRow(sensor Sensor, row int) *Interval {
	rowDist := mymath.Abs(row - sensor.Position.Y)

	if rowDist > sensor.Distance {
		// Out of range
		return nil
	}

	n := sensor.Distance - rowDist
	from := sensor.Position.X - n
	to := sensor.Position.X + n

	if from > p.max || to < 0 {
		// Out of range
		return nil
	}

	// limit
	if from < 0 {
		from = 0
	}
	if to > p.max {
		to = p.max
	}

	return &Interval{From: from, To: to}
}
