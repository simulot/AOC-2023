package main

import (
	"AOC2023/internal"
	"bytes"
	"fmt"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("noteSummary ", noteSummary(readPatterns(lines)))
}

type pattern [][]byte

func readPatterns(lines []string) []pattern {
	ps := []pattern{}
	p := pattern{}
	for _, l := range lines {
		if len(l) == 0 {
			if len(p) > 0 {
				ps = append(ps, p)
			}
			p = pattern{}
			continue
		}
		p = append(p, []byte(l))
	}
	if len(p) > 0 {
		ps = append(ps, p)
	}
	return ps
}

func horizontalAxis(p pattern) int {
	for r := 0; r < len(p)-1; r++ {
		isSymmetrical := true
		ra, rb := r, r+1
		for m := r; m < len(p); m++ {
			if !bytes.Equal(p[ra], p[rb]) {
				isSymmetrical = false
				break
			}
			ra--
			rb++
			if ra < 0 || rb >= len(p) {
				break
			}
		}
		if isSymmetrical {
			return r + 1
		}
	}
	return 0
}

func transposePattern(p pattern) pattern {
	t := make(pattern, len(p[0]))
	for c := range p[0] {
		v := make([]byte, len(p))
		for r := range p {
			v[r] = p[r][c]
		}
		t[c] = v
	}
	return t
}

func verticalAxis(p pattern) int {
	return horizontalAxis(transposePattern(p))
}

func noteSummary(ps []pattern) int {
	sum := 0
	for _, p := range ps {
		h := horizontalAxis(p)
		v := verticalAxis(p)
		sum += 100*h + v
	}
	return sum
}
