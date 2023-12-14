package main

import (
	"AOC2023/internal"
	"fmt"
	"hash/maphash"
	"slices"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("weightNorthBeam = ", cyclesAndWeight(readPlatform(lines), 1000000000))
}

type platform [][]byte

func readPlatform(lines []string) platform {
	p := platform{}
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		p = append(p, []byte(l))
	}
	return p
}

func cyclesAndWeight(p platform, n int) int {
	hashes := []uint64{}
	weights := []int{}

	iteration := 0
	var cycleBegin, cycleEnd int

	for iteration < n {
		h := getHash(p)
		if cycleBegin = slices.IndexFunc(hashes, func(e uint64) bool { return e == h }); cycleBegin >= 0 {
			cycleEnd = iteration
			break
		}
		weights = append(weights, weightNorthBeam(p))
		hashes = append(hashes, h)
		p = cycle(p)
		iteration++
		if iteration%10000 == 0 {
			fmt.Printf("Cycle %d\n", iteration)
		}
	}

	fmt.Printf("Cycle start at %d, length: %d\n", cycleBegin, cycleEnd)
	i := cycleBegin + (n-cycleBegin)%(cycleEnd-cycleBegin)
	return weights[i]
}

var hash maphash.Hash

func getHash(p platform) uint64 {
	hash.Reset()
	for r := range p {
		hash.Write(p[r])
	}
	return hash.Sum64()
}

func deepCopy(p platform) platform {
	np := make([][]byte, len(p))
	for r := range p {
		copy(np[r], p[r])
	}
	return np
}

func cycle(p platform) platform {
	p = tiltNorth(p)
	p = tiltWest(p)
	p = tiltSouth(p)
	p = tiltEast(p)
	return p
}

func tiltNorth(p platform) platform {
	done := false
	for !done {
		done = true
		for r := 1; r < len(p); r++ {
			for c, v := range p[r] {
				if v == 'O' && p[r-1][c] == '.' {
					p[r-1][c] = 'O'
					p[r][c] = '.'
					done = false
				}
			}
		}
	}
	return p
}

func tiltSouth(p platform) platform {
	done := false
	for !done {
		done = true
		for r := len(p) - 2; r >= 0; r-- {
			for c, v := range p[r] {
				if v == 'O' && p[r+1][c] == '.' {
					p[r+1][c] = 'O'
					p[r][c] = '.'
					done = false
				}
			}
		}
	}
	return p
}

func tiltWest(p platform) platform {
	done := false
	for !done {
		done = true
		for r := 0; r < len(p); r++ {
			for c := 1; c < len(p[r]); c++ {
				if p[r][c] == 'O' && p[r][c-1] == '.' {
					p[r][c-1] = 'O'
					p[r][c] = '.'
					done = false
				}
			}
		}
	}
	return p
}

func tiltEast(p platform) platform {
	done := false
	for !done {
		done = true
		for r := 0; r < len(p); r++ {
			for c := len(p[r]) - 2; c >= 0; c-- {
				if p[r][c] == 'O' && p[r][c+1] == '.' {
					p[r][c+1] = 'O'
					p[r][c] = '.'
					done = false
				}
			}
		}
	}
	return p
}

func weightNorthBeam(p platform) int {
	weight := 0
	for r := range p {
		for _, v := range p[r] {
			if v == 'O' {
				weight += len(p) - r
			}
		}
	}
	return weight
}
