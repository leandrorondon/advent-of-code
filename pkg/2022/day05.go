package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day05(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var stacks [9]Stack
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line[0] != '[' {
			continue
		}

		i := 0
		for c := 1; c < len(line); c = c + 4 {
			if line[c] != ' ' {
				stacks[i].Prepend(string(line[c]))
			}
			i++
		}
	}

	stacks2 := stacks

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		s := strings.Split(line, " ")
		n, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		for i := 0; i < n; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}

		stacks2[to-1].PushMultiple(stacks2[from-1].PopMultiple(n))
	}

	took := time.Now().Sub(t)
	var s, s2 string
	for i := 0; i < 9; i++ {
		s = s + stacks[i].Top()
		s2 = s2 + stacks2[i].Top()
	}

	fmt.Println("Crates on top:", s)
	fmt.Println("Crates on top 2:", s2)
	fmt.Printf("(took %v)\n", took)

	return nil
}

type Stack []string

func (s *Stack) Push(b string) {
	*s = append(*s, b)
}

func (s *Stack) Prepend(b string) {
	*s = append([]string{b}, *s...)
}

func (s *Stack) Pop() string {
	n := len(*s) - 1
	b := (*s)[n]
	*s = (*s)[:n]
	return b
}

func (s *Stack) Top() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) PopMultiple(n int) []string {
	size := len(*s)
	m := (*s)[size-n : size]
	*s = (*s)[:size-n]
	return m
}

func (s *Stack) PushMultiple(m []string) {
	*s = append(*s, m...)
}
