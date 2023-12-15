package main

import (
	"AOC2023/internal"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	b, err := os.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("focusingPower= ", focusingPower(readInstructions(internal.CSVSplit(strings.TrimSuffix(string(b), "\n")))))
}

type lens struct {
	label string
	focal int
}
type lenses map[byte][]lens

func readInstructions(instructions []string) lenses {
	lenses := lenses{}
	for _, ins := range instructions {
		newLens, op := decodeInstruction(ins)
		boxId := byte(hash(newLens.label))
		switch op {
		case OpRemove:
			if i := slices.IndexFunc[[]lens](lenses[boxId], func(l lens) bool {
				return newLens.label == l.label
			}); i >= 0 {
				lenses[boxId] = slices.Delete(lenses[boxId], i, i+1)
			}
			if len(lenses[boxId]) == 0 {
				delete(lenses, boxId)
			}
		case OpAdd:
			if i := slices.IndexFunc[[]lens](lenses[boxId], func(l lens) bool {
				return newLens.label == l.label
			}); i >= 0 {
				lenses[boxId][i] = newLens
			} else {
				lenses[boxId] = append(lenses[boxId], newLens)
			}
		}
	}
	return lenses
}

type operation byte

const (
	OpRemove operation = '-'
	OpAdd              = '='
)

func decodeInstruction(i string) (lens, operation) {
	var op operation
	var focal int
	var label string

	if strings.HasSuffix(i, "-") {
		op = '-'
		label = i[:len(i)-1]
	} else {
		op = operation(i[len(i)-2])
		focal = int(i[len(i)-1] - '0')
		label = i[:len(i)-2]
	}
	return lens{label: label, focal: focal}, op
}

func focusingPower(wall lenses) int {
	power := 0
	for b, ls := range wall {
		for slot, l := range ls {
			p := (int(b) + 1) * (slot + 1) * l.focal
			power += p
		}
	}
	return power
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
