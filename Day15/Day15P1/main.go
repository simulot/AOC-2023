package main

import (
	"AOC2023/internal"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sumOfHashes= ", sumOfHashes(internal.CSVSplit(strings.TrimSuffix(string(b), "\n"))))
}

func hash(s string) int {
	v := 0
	for _, c := range s {
		v += int(c)
		v *= 17
		v = v % 256
	}
	return v
}

func sumOfHashes(ss []string) int {
	sum := 0
	for _, s := range ss {
		h := hash(s)
		fmt.Printf("%q = %d\n", s, h)
		sum += h
	}
	return sum
}
