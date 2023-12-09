package main

import (
	"AOC2023/internal"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nearest location: ", nearestSoil(readAlmanac(lines)))
}

type Destination struct {
	From, To, Length int
}

type Map []Destination

type Almanac struct {
	Seeds []int
	Maps  []Map
}

func readAlmanac(lines []string) Almanac {
	a := Almanac{
		Maps: make([]Map, 7),
	}
	mapIndex := -1
	for _, l := range lines {
		if a.Seeds == nil && len(l) > 0 {
			a.Seeds = internal.ScanNumbers(strings.TrimPrefix(l, "seeds: "))
			continue
		}
		if strings.HasSuffix(l, "map:") {
			mapIndex++
			continue
		}
		if len(l) > 0 {
			numbers := internal.ScanNumbers(l)
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
			seed = d.To + +seed - d.From
			break
		}
	}
	return seed
}

func nearestSoil(a Almanac) int {
	minSoil := math.MaxInt

	for _, seed := range a.Seeds {
		soil := seed
		for _, m := range a.Maps {
			soil = mapSeedToSoil(m, soil)
		}
		minSoil = min(minSoil, soil)
	}
	return minSoil
}
