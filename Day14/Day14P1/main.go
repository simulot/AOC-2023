package main

import (
	"AOC2023/internal"
	"fmt"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("weightNorthBeam = ", weightNorthBeam(tiltNorth(readPlatform(lines))))
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
