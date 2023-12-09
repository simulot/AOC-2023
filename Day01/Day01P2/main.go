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

	fmt.Println("Calibration: ", trebuchet(lines))
}

func trebuchet(lines []string) int {
	sum := 0
	for _, l := range lines {
		var first, last int
		for i := 0; i < len(l); i++ {
			if first == 0 && l[i] >= '0' && l[i] <= '9' {
				first = int(l[i]) - '0'
			}
			if first == 0 {
				first = checkIsNumber(l[i:])
			}
			if last == 0 && l[len(l)-i-1] >= '0' && l[len(l)-i-1] <= '9' {
				last = int(l[len(l)-i-1]) - '0'
			}
			if last == 0 {
				last = checkIsNumber(l[len(l)-i-1:])
			}

			if first != 0 && last != 0 {
				sum += first*10 + last
				break
			}
		}
	}
	return sum
}

func checkIsNumber(s string) int {
	for d, number := range numbers {
		if strings.HasPrefix(s, number) {
			return d
		}
	}
	return 0
}

var numbers = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}
