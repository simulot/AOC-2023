package main

import (
	"AOC2023/internal"
	"fmt"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sum of Card points: ", scratchCards(lines))
}

func scratchCards(lines []string) int {
	sum := 0
	for _, l := range lines {
		ss := strings.Split(l, ":")
		ss = strings.Split(ss[1], "|")
		winningNumber := scanNumber(ss[0])
		numbers := scanNumber(ss[1])

		wins := 0
		points := 0
		for n := range numbers {
			if _, ok := winningNumber[n]; ok {
				if wins == 0 {
					points = 1
					wins = 1
				} else {
					wins += 1
					points *= 2
				}
			}
		}

		sum += points
	}

	return sum
}

func scanNumber(s string) map[string]any {
	numbers := map[string]any{}
	nn := strings.Fields(s)
	for _, n := range nn {
		numbers[n] = nil
	}
	return numbers
}
