package main

import (
	"AOC2023/internal"
	"reflect"
	"testing"
)

func Test_horizontalAxis(t *testing.T) {
	tc := []struct {
		name    string
		pattern pattern
		want    int
	}{
		{
			name: "test1",
			pattern: pattern{
				[]byte("#...##..#"),
				[]byte("#....#..#"),
				[]byte("..##..###"),
				[]byte("#####.##."),
				[]byte("#####.##."),
				[]byte("..##..###"),
				[]byte("#....#..#"),
			},
			want: 4,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := horizontalAxis(tt.pattern)
			if got != tt.want {
				t.Errorf("horizontalAxis()=%d, want:%d", got, tt.want)
			}
		})
	}
}
func Test_transposePattern(t *testing.T) {
	p := pattern{
		[]byte("#...##..#"),
		[]byte("#....#..#"),
		[]byte("..##..###"),
		[]byte("#####.##."),
		[]byte("#####.##."),
		[]byte("..##..###"),
		[]byte("#....#..#"),
	}
	want := pattern{
		[]byte("##.##.#"),
		[]byte("...##.."),
		[]byte("..####."),
		[]byte("..####."),
		[]byte("#..##.."),
		[]byte("##....#"),
		[]byte("..####."),
		[]byte("..####."),
		[]byte("###..##"),
	}
	got := transposePattern(p)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("transposePattern()")
	}

}

func Test_verticalAxis(t *testing.T) {
	tc := []struct {
		name    string
		pattern pattern
		want    int
	}{
		{
			name: "test1",
			pattern: pattern{
				[]byte("#.##..##."),
				[]byte("..#.##.#."),
				[]byte("##......#"),
				[]byte("##......#"),
				[]byte("..#.##.#."),
				[]byte("..##..##."),
				[]byte("#.#.##.#."),
			},
			want: 5,
		},

		{
			name: "test2",
			pattern: pattern{
				[]byte("..#....#."),
				[]byte("..######."),
				[]byte("..###.##."),
				[]byte("...####.."),
				[]byte("###.##.##"),
				[]byte("##.####.#"),
				[]byte("...#..#.."),
				[]byte("..######."),
			},
			want: 1,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := verticalAxis(tt.pattern)
			if got != tt.want {
				t.Errorf("horizontalAxis()=%d, want:%d", got, tt.want)
			}
		})
	}
}
func Test_noteSummary(t *testing.T) {
	tc := []struct {
		name     string
		patterns []pattern
		want     int
	}{
		{
			name: "set1",
			patterns: readPatterns(internal.ReadString(`
			#.##..##.
			..#.##.#.
			##......#
			##......#
			..#.##.#.
			..##..##.
			#.#.##.#.

			#...##..#
			#....#..#
			..##..###
			#####.##.
			#####.##.
			..##..###
			#....#..#`)),
			want: 405,
		},
		{
			name: "set2",
			patterns: readPatterns(internal.ReadString(`
			..##.
			##..#
			#...#
			..##.
			..##.
			`)),
			want: 400,
		},
		{
			name: "set3",
			patterns: readPatterns(internal.ReadString(`
			..##.
			..##.
			#...#
			##.##.
			..##.
			`)),
			want: 100,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := noteSummary(tt.patterns)
			if got != tt.want {
				t.Errorf("noteSummary()=%d, want:%d", got, tt.want)
			}
		})
	}
}
