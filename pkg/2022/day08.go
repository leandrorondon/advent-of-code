package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Day08(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var trees [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		trees = append(trees, []byte(scanner.Text()))
	}

	visibleCount := 0
	gridSize := len(trees)
	visible := makeGrid[bool](gridSize)
	highestScenicScore := 0

	// Top and bottom lines of trees are visible.
	heightsVisibleFromTop := trees[0]
	heightsVisibleFromBottom := trees[gridSize-1]
	visibleCount += 4*gridSize - 4

	n := gridSize - 1
	for x := 1; x < n; x++ {
		// Tree on the left and right are visible.
		heightVisibleFromLeft := trees[x][0]
		heightVisibleFromRight := trees[n-x][n]

		for y := 1; y < n; y++ {
			score := scenicScore(trees, x, y)
			if score > highestScenicScore {
				highestScenicScore = score
			}

			// Visibility from the top
			evaluate(trees[x][y], &heightsVisibleFromTop[y], &visible[x][y], &visibleCount)

			// Visibility from the left
			evaluate(trees[x][y], &heightVisibleFromLeft, &visible[x][y], &visibleCount)

			// Visibility from the bottom
			evaluate(trees[n-x][n-y], &heightsVisibleFromBottom[n-y], &visible[n-x][n-y], &visibleCount)

			// Visibility from the right
			evaluate(trees[n-x][n-y], &heightVisibleFromRight, &visible[n-x][n-y], &visibleCount)
		}
	}

	took := time.Now().Sub(t)
	fmt.Println("Trees visible from outside the grid:", visibleCount)
	fmt.Println("Highest scenic score:", highestScenicScore)
	fmt.Printf("(took %v)\n", took)

	return nil
}

func evaluate(tree byte, maxVisible *byte, visible *bool, count *int) {
	if tree > *maxVisible {
		*maxVisible = tree
		if !*visible {
			*visible = true
			*count++
		}
	}
}

func makeGrid[T any](size int) [][]T {
	grid := make([][]T, size)
	for i := range grid {
		grid[i] = make([]T, size)
	}
	return grid
}

func scenicScore(g [][]byte, x, y int) int {
	var left, right, up, down int

	v := g[x][y]

	for i := y + 1; i < len(g); i++ {
		right++
		if g[x][i] >= v {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		left++
		if g[x][i] >= v {
			break
		}
	}

	for i := x + 1; i < len(g); i++ {
		down++
		if g[i][y] >= v {
			break
		}

	}

	for i := x - 1; i >= 0; i-- {
		up++
		if g[i][y] >= v {
			break
		}

	}

	return left * right * up * down
}
