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

	fmt.Println("Total values: ", readAllValues(lines))
}

func readAllValues(lines []string) int {
	sum := 0
	for _, l := range lines {
		if l != "" {
			sum += processHistory(internal.ScanNumbers(l))
		}
	}
	return sum
}

func processHistory(v []int) int {
	v = append([]int{0}, v...)
	derivation := [][]int{v}
	row := 0
	for {
		next := make([]int, len(derivation[row])-1)
		allZeros := true
		for i := range derivation[row][0 : len(derivation[row])-1] {
			if i == 0 {
				continue
			}
			d := derivation[row][i+1] - derivation[row][i]
			next[i] = d
			allZeros = allZeros && d == 0
		}
		derivation = append(derivation, next)
		row++
		if allZeros {
			break
		}
	}
	if len(derivation) == 1 {
		panic("len(derivation) == 1")
	}
	value := 0
	for row := len(derivation) - 1; row > 0; row-- {
		a := derivation[row-1][1]
		b := derivation[row][0]
		value = a - b
		derivation[row-1][0] = value
	}
	return value
}
