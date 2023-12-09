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

	fmt.Println("Total steps: ", getSteps(readMap(lines)))
}

type Map struct {
	directions string
	nodes      map[string]node
}

type node struct {
	left, right string
}

func readMap(lines []string) Map {
	m := Map{
		nodes: make(map[string]node),
	}
	for _, l := range lines {
		if len(l) > 0 {
			if len(m.directions) == 0 {
				m.directions = l
				continue
			}
			var s node
			var at string
			fmt.Sscanf(l, "%3s = (%3s, %3s)", &at, &s.left, &s.right)
			m.nodes[at] = s
		}
	}
	return m
}

func getSteps(m Map) int {
	node := "AAA"
	instruction := 0
	steps := 0
	for node != "ZZZ" {
		turn := m.directions[instruction]
		s := m.nodes[node]
		switch turn {
		case 'L':
			node = s.left
		case 'R':
			node = s.right
		}
		instruction = (instruction + 1) % len(m.directions)
		steps++
	}
	return steps
}
