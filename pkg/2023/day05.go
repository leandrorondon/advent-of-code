package puzzles2023

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"
	
	"github.com/leandrorondon/advent-of-code/pkg/2023/day05"
)

func Day05(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	seeds, maps := day05.ParseLines(string(b))

	lowest := int64(math.MaxInt64)
	for _, seed := range seeds {
		location := maps.GetLocation(seed)
		if location < lowest {
			lowest = location
		}
	}

	lowest2 := int64(math.MaxInt64)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for idx := 0; idx < len(seeds); idx += 2 {
		wg.Add(1)
		go func(i int) {
			low := lowest2
			for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
				location := maps.GetLocation(seed)
				if location < low {
					low = location
				}
			}
			mutex.Lock()
			if low < lowest2 {
				lowest2 = low
			}
			mutex.Unlock()
			wg.Done()
		}(idx)
	}

	wg.Wait()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", lowest)
	fmt.Println("Part 2:", lowest2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
