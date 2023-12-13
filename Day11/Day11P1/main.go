package main

import (
	"AOC2023/internal"
	"fmt"
	"slices"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum Of Shortest Distances ", getSumOfShortestDistances(searchGalaxies(expand(readUniverse(lines)))))
}

type universe struct {
	image image
}
type position struct {
	horizontal, vertical int
}

func (u universe) Strings() []string {
	ret := []string{}
	for r := range u.image {
		sb := strings.Builder{}
		for _, c := range u.image[r] {
			sb.WriteByte(byte(c))
		}
		ret = append(ret, sb.String())
	}
	return ret
}

type image [][]byte

func (i image) At(horizontal, vertical int) byte {
	return i[horizontal][vertical]
}

func readUniverse(lines []string) universe {
	lines = internal.FilterOutEmptyLines(lines)
	size := len(lines)
	u := universe{
		image: make(image, size),
	}
	for row, l := range lines {
		r := make([]byte, len(l))
		for c := range l {
			r[c] = l[c]
		}
		u.image[row] = r
	}
	return u
}

func expand(u universe) universe {
	// add line in between empty lines
	newImage := image{}
	for r := range u.image {
		newImage = append(newImage, u.image[r])
		if p := slices.Index(u.image[r], '#'); p < 0 {
			l := make([]byte, len(u.image[r]))
			for i := range l {
				l[i] = '.'
			}
			newImage = append(newImage, l)
		}
	}

	// add a column between empty columns
	for c := len(newImage[0]) - 1; c >= 0; c-- {
		isEmpty := true
		for r := range newImage {
			if newImage.At(r, c) != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			for r := range newImage {
				newImage[r] = slices.Insert(newImage[r], c, '.')
			}
		}
	}

	return universe{
		image: newImage,
	}
}

func searchGalaxies(u universe) []position {
	galaxies := []position{}
	for v := range u.image {
		for h, c := range u.image[v] {
			if c == '#' {
				galaxies = append(galaxies, position{horizontal: h, vertical: v})
			}
		}
	}
	return galaxies
}

func getDistance(galaxies []position, a, b int) int {
	dh := abs(galaxies[a].horizontal - galaxies[b].horizontal)
	dv := abs(galaxies[a].vertical - galaxies[b].vertical)
	return dh + dv
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getSumOfShortestDistances(galaxies []position) int {
	sum := 0
	for g1 := range galaxies[:len(galaxies)-1] {
		for g2 := range galaxies[g1+1:] {
			a, b := g1, g1+1+g2
			d := getDistance(galaxies, a, b)
			sum += d
		}
	}
	return sum
}
