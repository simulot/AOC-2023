package main

import (
	"AOC2023/internal"
	"fmt"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Calibration: ", trebuchet(lines))
}

func trebuchet(lines []string) int {
	sum := 0
	for _, l := range lines {
		var first, last byte
		for i := 0; i < len(l); i++ {
			if first == 0 && l[i] >= '0' && l[i] <= '9' {
				first = l[i]
			}
			if last == 0 && l[len(l)-i-1] >= '0' && l[len(l)-i-1] <= '9' {
				last = l[len(l)-i-1]
			}
			if first != 0 && last != 0 {
				sum += int((first-'0')*10) + int((last - '0'))
				break
			}
		}
	}
	return sum
}
