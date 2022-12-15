package grid

import (
	"sort"

	mymath "github.com/leandrorondon/advent-of-code/pkg/math"
)

type Coordinate struct {
	X int
	Y int
}

type SensonBeaconPair struct {
	Sensor Coordinate
	Beacon Coordinate
}

type Item string

type Sensor struct {
	Position      Coordinate
	ClosestBeacon Coordinate
	Distance      int
}

type Interval struct {
	From int
	To   int
}

func New(pairs []SensonBeaconPair) *Grid {
	g := &Grid{}

	for _, pair := range pairs {
		beacon := pair.Beacon
		sensor := pair.Sensor

		dist := g.ManhattanDistance(beacon, sensor)
		g.sensors = append(g.sensors, Sensor{
			Position:      sensor,
			ClosestBeacon: beacon,
			Distance:      dist,
		})
	}

	return g
}

type Grid struct {
	sensors []Sensor
}

func (g *Grid) ManhattanDistance(c1, c2 Coordinate) int {
	return mymath.Abs(c1.X-c2.X) + mymath.Abs(c1.Y-c2.Y)
}

func (g *Grid) CountNoBeaconInRow(row int) int {
	var intervals []Interval
	for _, sensor := range g.sensors {
		interval := g.ProjectInRow(sensor, row)
		if interval != nil {
			intervals = append(intervals, *interval)
		}
	}

	merged := g.MergeIntervals(intervals)
	return g.CountSlots(merged)
}

func (g *Grid) CountSlots(intervals []Interval) int {
	count := 0
	for _, iv := range intervals {
		count += iv.To - iv.From
	}
	return count
}

func (g *Grid) MergeIntervals(orig []Interval) []Interval {
	if len(orig) == 0 {
		return orig
	}

	sort.Slice(orig, func(i, j int) bool {
		return orig[i].From < orig[j].From
	})

	var intervals []Interval

	count := 1
	current := orig[0]
	for i := 1; i < len(orig); i++ {
		if orig[i].From <= current.To { // intersection
			if orig[i].To > current.To { // partial intersection
				current.To = orig[i].To
			} // else: orig[i] contains orig[i-1]
			continue
		}

		// no intersection
		intervals = append(intervals, current)
		current = orig[i]
		count++
	}
	if count > len(intervals) {
		intervals = append(intervals, current)
	}

	return intervals
}

func (g *Grid) ProjectInRow(sensor Sensor, row int) *Interval {
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
