package main

import (
	"AOC2023/internal"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Nearest location: ", nearestSoil(readAlmanac(lines)))
}

type Destination struct {
	From, To, Length int
}

type Map []Destination

type SeedRange struct {
	start  int
	number int
}
type Almanac struct {
	Seeds []SeedRange
	Maps  []Map
}

func readAlmanac(lines []string) Almanac {
	a := Almanac{
		Maps: make([]Map, 7),
	}
	mapIndex := -1
	for _, l := range lines {
		if a.Seeds == nil && len(l) > 0 {
			n := scanNumbers(strings.TrimPrefix(l, "seeds: "))
			a.Seeds = make([]SeedRange, len(n)/2)
			for i := 0; i < len(n); i += 2 {
				a.Seeds[i/2].start = n[i]
				a.Seeds[i/2].number = n[i+1]
			}
			continue
		}
		if strings.HasSuffix(l, "map:") {
			mapIndex++
			continue
		}
		if len(l) > 0 {
			numbers := scanNumbers(l)
			if len(numbers) != 3 {
				panic("incorrect map")
			}
			m := a.Maps[mapIndex]
			m = append(m, Destination{To: numbers[0], From: numbers[1], Length: numbers[2]})
			a.Maps[mapIndex] = m
		}
	}
	return a
}

func mapSeedToSoil(m Map, seed int) int {
	for _, d := range m {
		if seed >= d.From && seed < d.From+d.Length {
			seed = d.To + seed - d.From
			break
		}
	}
	return seed
}

func nearestSoil(a Almanac) int {
	minSoil := math.MaxInt
	l := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, seedRange := range a.Seeds {
		wg.Add(1)
		go func(seedRange SeedRange) {
			for seed := seedRange.start; seed < seedRange.start+seedRange.number; seed++ {
				soil := seed
				for _, m := range a.Maps {
					soil = mapSeedToSoil(m, soil)
				}
				l.Lock()
				minSoil = min(minSoil, soil)
				l.Unlock()
			}
			wg.Done()
		}(seedRange)
	}
	wg.Wait()
	return minSoil
}

func scanNumbers(l string) []int {
	ss := strings.Fields(l)
	numbers := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers[i] = n
	}
	return numbers
}
