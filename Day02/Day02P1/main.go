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
		return
	}

	fmt.Println("Sum of valid games: ", validGames(lines, map[color]int{
		red:   12,
		green: 13,
		blue:  14,
	},
	))
}

func validGames(lines []string, bag map[color]int) int {
	sum := 0
	for _, l := range lines {
		ss := strings.Split(l, ":")
		var game int
		fmt.Sscanf(ss[0], "Game %d", &game)
		isValid := true

		ss = strings.Split(ss[1], ";")
		for _, s := range ss {
			cc := strings.Split(s, ",")
			for _, c := range cc {
				var col color
				var n int
				fmt.Sscanf(c, "%d %s", &n, &col)
				if n > bag[col] {
					isValid = false
				}
			}
		}
		if isValid {
			sum += game
		}
	}

	return sum
}

type color string

const (
	blue  color = "blue"
	red         = "red"
	green       = "green"
)
