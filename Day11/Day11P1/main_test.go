package main

import (
	"AOC2023/internal"
	"fmt"
	"reflect"
	"testing"
)

func Test_expand(t *testing.T) {
	tc := []struct {
		name  string
		lines []string
		want  []string
	}{
		{
			name: "example",
			lines: internal.ReadString(`
			...#......
			.......#..
			#.........
			..........
			......#...
			.#........
			.........#
			..........
			.......#..
			#...#.....`),
			want: internal.FilterOutEmptyLines(internal.ReadString(`
			....#........
			.........#...
			#............
			.............
			.............
			........#....
			.#...........
			............#
			.............
			.............
			.........#...
			#....#.......
			`)),
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := expand(readUniverse(tt.lines)).Strings()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expand() lines %d,%d", len(got), len(tt.want))
				for r := 0; r < min(len(got), len(tt.want)); r++ {
					t.Error(got[r], "  ", tt.want[r])
				}
			}
		})
	}
}

func Test_getDistance(t *testing.T) {
	u := readUniverse(internal.FilterOutEmptyLines(internal.ReadString(`
	....#........
	.........#...
	#............
	.............
	.............
	........#....
	.#...........
	............#
	.............
	.............
	.........#...
	#....#.......
	`)))
	g := searchGalaxies(u)
	tc := []struct {
		couple [2]int
		want   int
	}{
		{
			couple: [2]int{1, 7},
			want:   15,
		},
		{
			couple: [2]int{3, 6},
			want:   17,
		},
		{
			couple: [2]int{8, 9},
			want:   5,
		},
	}

	for _, tt := range tc {
		t.Run(fmt.Sprintf("(%d,%d)", tt.couple[0], tt.couple[1]), func(t *testing.T) {
			got := getDistance(g, tt.couple[0]-1, tt.couple[1]-1)
			if got != tt.want {
				t.Errorf("getDistance(%d,%d)=%d, want %d", tt.couple[0], tt.couple[1], got, tt.want)
			}
		})
	}
}

func Test_getSumOfShortestDistances(t *testing.T) {
	tc := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "example",
			lines: internal.ReadString(`
			...#......
			.......#..
			#.........
			..........
			......#...
			.#........
			.........#
			..........
			.......#..
			#...#.....`),
			want: 374,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := getSumOfShortestDistances(searchGalaxies(expand(readUniverse(tt.lines))))
			if tt.want != got {
				t.Errorf("getSumOfShortestDistances()=%d,%d", got, tt.want)
			}
		})
	}
}
