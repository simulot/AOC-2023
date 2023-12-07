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

	fmt.Println("Win margin: ", winMargin(readRaces(lines)))
}

type race struct {
	t int
	d int
}

func readRaces(lines []string) []race {
	var (
		times     []int
		distances []int
	)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		ss := strings.Split(l, ":")
		switch {
		case strings.HasPrefix(ss[0], "Time"):
			s := strings.ReplaceAll(ss[1], " ", "")
			times = internal.ScanNumbers(s)
		case strings.HasPrefix(ss[0], "Distance"):
			s := strings.ReplaceAll(ss[1], " ", "")
			distances = internal.ScanNumbers(s)
		}
	}

	races := []race{}
	for i := range times {
		races = append(races, race{t: times[i], d: distances[i]})
	}
	return races
}

func winMargin(races []race) int {
	margin := 1
	for _, r := range races {
		margin *= checkRace(r)
	}
	return margin
}

func checkRace(r race) int {
	count := 0
	for i := 0; i <= r.t; i++ {
		d := i * (r.t - i)
		if d < 0 {
			break
		}
		if d > r.d {
			count++
		}
	}
	return count
}
