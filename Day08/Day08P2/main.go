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
	steps := []int{}
	for s := range m.nodes {
		if strings.HasSuffix(s, "A") {
			steps = append(steps, getStepsForNode(m, s))
		}
	}
	return lcm(steps...)
}
func getStepsForNode(m Map, node string) int {
	instruction := 0
	steps := 0
	for !strings.HasSuffix(node, "Z") {
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

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// least common multiple
func lcm(n ...int) int {
	if len(n) < 2 {
		panic("lcm...")
	}
	r := n[0] * n[1] / gcd(n[0], n[1])
	for _, i := range n[2:] {
		r = lcm(r, i)
	}
	return r
}
