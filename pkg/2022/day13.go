package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day13/decoder"
)

func Day13(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := decoder.NewParser(scanner)
	signal := p.ParseSignal()

	t1 := time.Now()
	part1 := signal.PairsInRightOrder()
	took1 := time.Now().Sub(t1)

	t2 := time.Now()
	divider1 := p.ParsePacket("[[2]]")
	divider2 := p.ParsePacket("[[6]]")
	decoderKey := signal.DecoderKey(divider1, divider2)
	took2 := time.Now().Sub(t2)

	took := time.Now().Sub(t)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)
	fmt.Printf("- Part 2: %d (took %v)\n", decoderKey, took2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
