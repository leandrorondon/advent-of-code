package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var priority map[byte]int

func init() {
	priority = make(map[byte]int)
	for i := 0; i < 26; i++ {
		priority[byte('a'+i)] = i + 1
		priority[byte('A'+i)] = i + 27
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		m := make(map[byte]bool)
		for i := 0; i < len(line)/2; i++ {
			m[line[i]] = true
		}

		for i := len(line) / 2; i < len(line); i++ {
			c := line[i]
			if m[c] {
				sum += priority[c]
				break
			}
		}
	}

	fmt.Println("Sum 1:", sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var sum, count int
	var itemCount map[uint8]int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if count%3 == 0 {
			// new group, reset count

			itemCount = make(map[uint8]int)
		}
		count++

		rucksack := make(map[uint8]bool)
		for i := 0; i < len(line); i++ {
			item := line[i]
			if !rucksack[item] {
				rucksack[item] = true
				itemCount[item] = itemCount[item] + 1
				if itemCount[item] == 3 {
					// Badge found
					sum += priority[item]
					break
				}
			}
		}
	}

	fmt.Println("Sum 2:", sum)
}
