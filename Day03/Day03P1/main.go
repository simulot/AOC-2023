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

	fmt.Println("Sum of part number: ", parNumbers(lines))
}

var getNumbers = regexp.MustCompile(`\d+`)

func parNumbers(lines []string) int {
	sum := 0
	for row, l := range lines {
		numbers := getNumbers.FindAllStringIndex(l, -1)
		for _, n := range numbers {
			isPart := false
			if !isPart && n[0] > 0 {
				isPart = isSymbol(l[n[0]-1])
			}
			if !isPart && n[1] < len(l) {
				isPart = isSymbol(l[n[1]])
			}
			if !isPart && row > 0 {
				for c := max(0, n[0]-1); c <= min(n[1], len(l)-1); c++ {
					if isSymbol(lines[row-1][c]) {
						isPart = true
						break
					}
				}
			}
			if !isPart && row < len(lines)-1 {
				for c := max(0, n[0]-1); c <= min(n[1], len(l)-1); c++ {
					if isSymbol(lines[row+1][c]) {
						isPart = true
						break
					}
				}
			}
			if isPart {
				p, _ := strconv.Atoi(l[n[0]:n[1]])
				sum += p
			}
		}
	}

	return sum
}

func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}
