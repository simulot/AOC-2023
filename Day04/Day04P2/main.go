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

	fmt.Println("Number of Cards: ", scratchCardNumber(readCards(lines)))
}

type card struct {
	matchingNumbers int
	instances       int
}

func readCards(lines []string) []card {
	cards := make([]card, len(lines))
	for c, l := range lines {
		ss := strings.Split(l, ":")
		ss = strings.Split(ss[1], "|")
		winningNumber := scanNumber(ss[0])
		numbers := scanNumber(ss[1])

		matching := 0
		for n := range numbers {
			if _, ok := winningNumber[n]; ok {
				matching++
			}
		}
		cards[c] = card{matchingNumbers: matching, instances: 1}
	}
	return cards
}

func scratchCardNumber(cards []card) int {
	for i := 0; i < len(cards); i++ {
		for k := cards[i].instances; k > 0; k-- {
			got := cards[i].matchingNumbers
			for j := i + 1; j < min(i+got+1, len(cards)); j++ {
				cards[j].instances++
			}
		}
	}
	sum := 0
	for i := range cards {
		sum += cards[i].instances
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
