package puzzles2023

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Day01(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")

	sum := 0

	for _, line := range lines {
		s := make([]byte, 2)

		l := len(line) - 1

		for i := range line {
			if s[0] == 0 && unicode.IsDigit(rune(line[i])) {
				s[0] = line[i]
			}

			if s[1] == 0 && unicode.IsDigit(rune(line[l-i])) {
				s[1] = line[l-i]
			}

			if s[0] != 0 && s[1] != 0 {
				break
			}
		}

		v, _ := strconv.Atoi(string(s))
		sum += v
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", sum)
	fmt.Printf("(took %v)\n", took)

	sum = 0

	for _, ll := range lines {

		line := replaceSpelledOutDigits(ll)

		s := make([]byte, 2)

		l := len(line) - 1

		for i := range line {
			if s[0] == 0 && unicode.IsDigit(rune(line[i])) {
				s[0] = line[i]
			}

			if s[1] == 0 && unicode.IsDigit(rune(line[l-i])) {
				s[1] = line[l-i]
			}

			if s[0] != 0 && s[1] != 0 {
				break
			}
		}
		v, _ := strconv.Atoi(string(s))
		sum += v
	}

	took = time.Now().Sub(t)
	fmt.Println("Part 2:", sum)
	fmt.Printf("(took %v)\n", took)

	return nil
}

func replaceSpelledOutDigits(s string) string {
	m := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for o, n := range m {
		s = strings.Replace(s, o, n, -1)
	}

	return s
}
