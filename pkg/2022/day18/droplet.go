package day18

import (
	"math"
	"sort"

	mymath "github.com/leandrorondon/advent-of-code/pkg/math"
)

type XYZ struct {
	X int
	Y int
	Z int
}

func (xyz XYZ) Sum(other XYZ) XYZ {
	return XYZ{
		xyz.X + other.X,
		xyz.Y + other.Y,
		xyz.Z + other.Z,
	}
}

type MinMax struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func SurfaceArea(droplets []XYZ) int {
	totalSides := 6 * len(droplets)

	sortXYZ(droplets)
	for i := 0; i < len(droplets)-1; i++ {
		if droplets[i].X != droplets[i+1].X || droplets[i].Y != droplets[i+1].Y {
			continue
		}

		if droplets[i].Z+1 == droplets[i+1].Z {
			totalSides -= 2
		}
	}

	sortXZY(droplets)
	for i := 0; i < len(droplets)-1; i++ {
		if droplets[i].X != droplets[i+1].X || droplets[i].Z != droplets[i+1].Z {
			continue
		}

		if droplets[i].Y+1 == droplets[i+1].Y {
			totalSides -= 2
		}
	}

	sortYZX(droplets)
	for i := 0; i < len(droplets)-1; i++ {
		if droplets[i].Y != droplets[i+1].Y || droplets[i].Z != droplets[i+1].Z {
			continue
		}

		if droplets[i].X+1 == droplets[i+1].X {
			totalSides -= 2
		}
	}

	return totalSides
}

func sortXYZ(droplets []XYZ) {
	sort.Slice(droplets, func(i, j int) bool {
		if droplets[i].X < droplets[j].X {
			return true
		}

		if droplets[i].X > droplets[j].X {
			return false
		}

		if droplets[i].Y < droplets[j].Y {
			return true
		}

		if droplets[i].Y > droplets[j].Y {
			return false
		}

		if droplets[i].Z < droplets[j].Z {
			return true
		}

		if droplets[i].Z > droplets[j].Z {
			return false
		}

		return true
	})
}

func sortXZY(droplets []XYZ) {
	sort.Slice(droplets, func(i, j int) bool {
		if droplets[i].X < droplets[j].X {
			return true
		}

		if droplets[i].X > droplets[j].X {
			return false
		}

		if droplets[i].Z < droplets[j].Z {
			return true
		}

		if droplets[i].Z > droplets[j].Z {
			return false
		}

		if droplets[i].Y < droplets[j].Y {
			return true
		}

		if droplets[i].Y > droplets[j].Y {
			return false
		}

		return true
	})
}

func sortYZX(droplets []XYZ) {
	sort.Slice(droplets, func(i, j int) bool {
		if droplets[i].Y < droplets[j].Y {
			return true
		}

		if droplets[i].Y > droplets[j].Y {
			return false
		}

		if droplets[i].Z < droplets[j].Z {
			return true
		}

		if droplets[i].Z > droplets[j].Z {
			return false
		}

		if droplets[i].X < droplets[j].X {
			return true
		}

		if droplets[i].X > droplets[j].X {
			return false
		}

		return true
	})
}

func ExternalSurface(droplets []XYZ) int {
	m, min, max := FindMapMinMax(droplets)
	min.X, min.Y, min.Z = min.X-1, min.Y-1, min.Z-1
	max.X, max.Y, max.Z = max.X+1, max.Y+1, max.Z+1
	cuboid := buildCuboid(m, min, max)
	cuboidExternalArea := 2*(max.X-min.X+1)*(max.Y-min.Y+1) + 2*(max.X-min.X+1)*(max.Z-min.Z+1) + 2*(max.Z-min.Z+1)*(max.Y-min.Y+1)
	suf := SurfaceArea(cuboid)
	return suf - cuboidExternalArea
}

var neighborsXY = []XYZ{
	{-1, 0, 0},
	{1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
}

var previousZDecr = XYZ{0, 0, 1}

var previousZCres = XYZ{0, 0, -1}

var previousYDecr = XYZ{0, 1, 0}

var previousYCres = XYZ{0, -1, 0}

var previousXDecr = XYZ{1, 0, 0}

var previousXCres = XYZ{-1, 0, 0}

var neighborsXZ = []XYZ{
	{-1, 0, 0},
	{1, 0, 0},
	{0, 0, 1},
	{0, 0, -1},
}

var neighborsYZ = []XYZ{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
}

func buildCuboid(m map[XYZ]bool, min, max XYZ) []XYZ {
	var cuboid []XYZ
	cuboidMap := make(map[XYZ]bool)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			// XY plan / z-axis / crescent
			for z := min.Z; z <= max.Z; z++ {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsXY, previousZCres) {
					break
				}
			}

			// XY plan / z-axis / decrescent
			for z := max.Z; z >= min.Z; z-- {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsXY, previousZDecr) {
					break
				}
			}
		}

		for z := min.Z; z <= max.Z; z++ {
			// XZ plan / y-axis / crescent
			for y := min.Y; y <= max.Y; y++ {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsXZ, previousYCres) {
					break
				}
			}

			// XZ plan / y-axis / decrescent
			for y := max.Y; y >= min.Y; y-- {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsXZ, previousYDecr) {
					break
				}
			}
		}
	}

	for y := min.Y; y <= max.Y; y++ {
		for z := min.Z; z <= max.Z; z++ {
			// YZ plan / x-axis / crescent
			for x := min.X; x <= max.X; x++ {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsYZ, previousXCres) {
					break
				}
			}

			// YZ plan / x-axis / decrescent
			for x := max.X; x >= min.X; x-- {
				if checkDroplet(&cuboid, m, cuboidMap, x, y, z, neighborsYZ, previousXDecr) {
					break
				}
			}
		}
	}

	return cuboid
}

func checkDroplet(cuboid *[]XYZ, dropletMap, cuboidMap map[XYZ]bool, x, y, z int, neighbors []XYZ, previous XYZ) bool {
	xyz := XYZ{x, y, z}

	// droplet found
	if dropletMap[xyz] {
		prev := xyz.Sum(previous)
		for _, neighbor := range neighbors {
			n := prev.Sum(neighbor)
			if dropletMap[n] {
				continue
			}

			checkAndFill(cuboid, cuboidMap, n)
		}

		return true
	}

	checkAndFill(cuboid, cuboidMap, xyz)

	return false
}

func checkAndFill(cuboid *[]XYZ, cuboidMap map[XYZ]bool, xyz XYZ) {
	// already in the cuboid
	if cuboidMap[xyz] {
		return
	}

	cuboidMap[xyz] = true
	*cuboid = append(*cuboid, xyz)
}

func FindMapMinMax(droplets []XYZ) (map[XYZ]bool, XYZ, XYZ) {
	min := XYZ{math.MaxInt, math.MaxInt, math.MaxInt}
	max := XYZ{math.MinInt, math.MinInt, math.MinInt}
	m := make(map[XYZ]bool)

	for _, d := range droplets {
		m[d] = true
		min.X = mymath.Min(min.X, d.X)
		min.Y = mymath.Min(min.Y, d.Y)
		min.Z = mymath.Min(min.Z, d.Z)
		max.X = mymath.Max(max.X, d.X)
		max.Y = mymath.Max(max.Y, d.Y)
		max.Z = mymath.Max(max.Z, d.Z)
	}

	return m, min, max
}
