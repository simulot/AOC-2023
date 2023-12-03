package main

import (
	"AOC2023/internal"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sum of gearRatio: ", gearRatio(lines))
}

var getNumbers = regexp.MustCompile(`\d+`)

func gearRatio(lines []string) int {
	sum := 0
	for row, l := range lines {
		for i := 0; i < len(l); i++ {
			multipliers := []int{}
			if l[i] != '*' {
				continue
			}
			for r := max(0, row-1); r <= min(row+1, len(lines)-1); r++ {
				numbers := getNumbers.FindAllStringIndex(lines[r], -1)
				for _, n := range numbers {
					isAdjacent := false
					switch {
					case r < row || r > row:
						isAdjacent = n[0]-1 <= i && i <= n[1]
					case r == row:
						isAdjacent = i == n[0]-1 || i == n[1]
					}
					if isAdjacent {
						multipliers = append(multipliers, getNumber(lines[r], n[0], n[1]))
					}
				}
			}
			if len(multipliers) == 2 {
				sum += multipliers[0] * multipliers[1]
			}
		}
	}
	return sum
}

func getNumber(s string, start, end int) int {
	i, err := strconv.Atoi(s[start:end])
	if err != nil {
		panic(err)
	}
	return i
}
