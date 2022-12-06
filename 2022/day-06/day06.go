package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	text, _ := os.ReadFile("input.txt")
	packet := detect(text, 4)
	message := detect(text, 14)
	took := time.Now().Sub(t)

	fmt.Println("Chars until first start-of-packet is detected:", packet)
	fmt.Println("Chars until first start-of-message is detected:", message)
	fmt.Printf("(took %v)\n", took)
}

func detect(text []byte, size int) int {
	m := make(map[byte]int)
	oldestIdx := 0

	for i, c := range text {
		m[c] = m[c] + 1
		if i < size {
			continue
		}

		oldC := text[oldestIdx]
		m[oldC] = m[oldC] - 1
		if m[oldC] == 0 {
			delete(m, oldC)
		}

		if len(m) == size {
			return i + 1
		}

		oldestIdx++
	}

	return 0
}
