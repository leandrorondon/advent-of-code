package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day07(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	root := NewDir("/")
	buildTree(root, scanner)

	sum := findSumDirsWithMaxSize(root, 100000)

	needsToFree := 30000000 - (70000000 - root.Size)
	size := SizeOfTheSmallestDirectoryThatIfDeletedWouldFreeUpEnoughSpaceOnTheFilesystemToRunTheUpdate(root, needsToFree)

	took := time.Now().Sub(t)

	fmt.Println("Sum part 1:", sum)
	fmt.Println("Size part 2:", size)
	fmt.Printf("(took %v)\n", took)

	return nil
}

func NewDir(name string) *Dir {
	return &Dir{
		Name:  name,
		Dirs:  make(map[string]*Dir),
		Files: make(map[string]int),
	}
}

type Dir struct {
	Dirs  map[string]*Dir
	Files map[string]int
	Name  string
	Size  int
}

func buildTree(dir *Dir, scanner *bufio.Scanner) string {
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")

		switch s[0] {
		case "$":
			// Command
			dest := command(dir, s[1:], scanner)
			if dest == ".." {
				return ""
			}
			if dest == dir.Name {
				continue
			}
		default:
			// Listing
			list(dir, s)
		}
	}

	return ""
}

func command(dir *Dir, s []string, scanner *bufio.Scanner) string {
	switch s[0] {
	case "cd":
		return cd(dir, s[1], scanner)
	default:
		// ls - ignore
	}

	return ""
}

func cd(dir *Dir, dest string, scanner *bufio.Scanner) string {
	switch dest {
	case "..", "/":
		return dest
	default:
		r := buildTree(dir.Dirs[dest], scanner)
		dir.Size += dir.Dirs[dest].Size
		return r
	}
}

func list(dir *Dir, s []string) {
	switch s[0] {
	case "dir":
		// Listing dir
		dir.Dirs[s[1]] = NewDir(s[1])
	default:
		// Listing file
		size, _ := strconv.Atoi(s[0])
		dir.Files[s[1]] = size
		dir.Size += size
	}
}

func findSumDirsWithMaxSize(dir *Dir, limit int) int {
	total := 0

	if dir.Size <= limit {
		total += dir.Size
	}

	for _, d := range dir.Dirs {
		total += findSumDirsWithMaxSize(d, limit)
	}

	return total
}

func printTree(dir *Dir, offset int) {
	o := strings.Repeat("  ", offset)
	fmt.Printf("%s - %s (%d)\n", o, dir.Name, dir.Size)
	for _, d := range dir.Dirs {
		printTree(d, offset+1)
	}
	for name, size := range dir.Files {
		fmt.Printf("%s | %s (%d)\n", o, name, size)
	}
}

func SizeOfTheSmallestDirectoryThatIfDeletedWouldFreeUpEnoughSpaceOnTheFilesystemToRunTheUpdate(dir *Dir, needsToFree int) int {
	smallest := 70000000

	if dir.Size >= needsToFree {
		smallest = dir.Size
	}

	for _, d := range dir.Dirs {
		found := SizeOfTheSmallestDirectoryThatIfDeletedWouldFreeUpEnoughSpaceOnTheFilesystemToRunTheUpdate(d, needsToFree)

		if found >= needsToFree && found < smallest {
			smallest = found
		}
	}

	return smallest
}
