package day18

import (
	"sort"
)

type Droplet struct {
	X int
	Y int
	Z int
}

func SurfaceArea(droplets []Droplet) int {
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

func sortXYZ(droplets []Droplet) {
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

func sortXZY(droplets []Droplet) {
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

func sortYZX(droplets []Droplet) {
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
