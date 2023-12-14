package main

import (
	"reflect"
	"testing"
)

func Test_tiltNorth(t *testing.T) {
	tc := []struct {
		name string
		p    platform
		want platform
	}{
		{
			name: "test1",
			p: platform{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			want: platform{
				[]byte("OOOO.#.O.."),
				[]byte("OO..#....#"),
				[]byte("OO..O##..O"),
				[]byte("O..#.OO..."),
				[]byte("........#."),
				[]byte("..#....#.#"),
				[]byte("..O..#.O.O"),
				[]byte("..O......."),
				[]byte("#....###.."),
				[]byte("#....#...."),
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := tiltNorth(tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tiltNorth()")
				for r := 0; r < min(len(got), len(tt.want)); r++ {
					t.Errorf("%-4d %s   %s", r, string(got[r]), string(tt.want[r]))
				}
			}
		})
	}
}

func Test_weightNorthBeam(t *testing.T) {
	tc := []struct {
		name string
		p    platform
		want int
	}{
		{
			name: "test1",
			p: platform{
				[]byte("OOOO.#.O.."),
				[]byte("OO..#....#"),
				[]byte("OO..O##..O"),
				[]byte("O..#.OO..."),
				[]byte("........#."),
				[]byte("..#....#.#"),
				[]byte("..O..#.O.O"),
				[]byte("..O......."),
				[]byte("#....###.."),
				[]byte("#....#...."),
			},
			want: 136,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := weightNorthBeam(tt.p)
			if got != tt.want {
				t.Errorf("weightNorthBeam()=%d, want:%d", got, tt.want)
			}
		})
	}
}
