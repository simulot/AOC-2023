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

	fmt.Println("Sum of power games: ", powerGames(lines))
}

type color string

func powerGames(lines []string) int {
	sum := 0
	for _, l := range lines {
		ss := strings.Split(l, ":")
		var game int
		fmt.Sscanf(ss[0], "Game %d", &game)
		maxBag := map[color]int{}
		ss = strings.Split(ss[1], ";")
		for _, s := range ss {
			cc := strings.Split(s, ",")
			for _, c := range cc {
				var col color
				var n int
				fmt.Sscanf(c, "%d %s", &n, &col)
				if max := maxBag[col]; n > max {
					maxBag[col] = n
				}
			}
		}
		power := 1
		for c := range maxBag {
			power = power * maxBag[c]
		}
		sum += power
	}

	return sum
}
