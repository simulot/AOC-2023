package main

import (
	"AOC2023/internal"
	"reflect"
	"testing"
)

func Test_lightIt(t *testing.T) {
	tc := []struct {
		name string
		c    *contraption
		want []string
	}{
		{
			name: "damn simple",
			c: readContraption(internal.FilterOutEmptyLines(internal.ReadString(`
			..\.
			.\/.`))),
			want: []string{
				"###.",
				".##.",
			},
		},

		{
			name: "sample",
			c: readContraption(internal.FilterOutEmptyLines(internal.ReadString(`
					.|...\....
					|.-.\.....
					.....|-...
					........|.
					..........
					.........\
					..../.\\..
					.-.-/..|..
					.|....-|.\
					..//.|....
				`))),
			want: []string{
				"######....",
				".#...#....",
				".#...#####",
				".#...##...",
				".#...##...",
				".#...##...",
				".#..####..",
				"########..",
				".#######..",
				".#...#.#..",
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.lightIt(beam{position: point{horizontal: 0, vertical: 0}, direction: '>'}).visualize()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference")
				for r := 0; r < min(len(tt.want), len(got)); r++ {
					t.Errorf("row %2d: %s  %s", r, got[r], tt.want[r])
				}
			}
		})
	}
}

func Test_energizedCount(t *testing.T) {
	tc := []struct {
		name string
		c    *contraption
		want int
	}{
		{
			name: "damn simple",
			c: readContraption(internal.FilterOutEmptyLines(internal.ReadString(`
			..\.
			.\/.`))),
			want: 5,
		},

		{
			name: "sample",
			c: readContraption(internal.FilterOutEmptyLines(internal.ReadString(`
					.|...\....
					|.-.\.....
					.....|-...
					........|.
					..........
					.........\
					..../.\\..
					.-.-/..|..
					.|....-|.\
					..//.|....
				`))),
			want: 46,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.lightIt(beam{position: point{horizontal: 0, vertical: 0}, direction: '>'}).energizedCount()
			if got != tt.want {
				t.Errorf("energizedCount=%d, want:%d", got, tt.want)
			}
		})
	}
}

func Test_maxEnergizedCount(t *testing.T) {
	tc := []struct {
		name string
		c    []string
		want int
	}{
		{
			name: "sample",
			c: internal.FilterOutEmptyLines(internal.ReadString(`
					.|...\....
					|.-.\.....
					.....|-...
					........|.
					..........
					.........\
					..../.\\..
					.-.-/..|..
					.|....-|.\
					..//.|....
				`)),
			want: 51,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := maxEnergizedCount(tt.c)
			if got != tt.want {
				t.Errorf("maxEnergizedCount=%d, want:%d", got, tt.want)
			}
		})
	}
}
