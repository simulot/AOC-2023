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

func Test_cycle(t *testing.T) {
	p := platform{
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
	}

	afterCycle := []platform{
		{
			[]byte(".....#...."),
			[]byte("....#...O#"),
			[]byte("...OO##..."),
			[]byte(".OO#......"),
			[]byte(".....OOO#."),
			[]byte(".O#...O#.#"),
			[]byte("....O#...."),
			[]byte("......OOOO"),
			[]byte("#...O###.."),
			[]byte("#..OO#...."),
		},
		{
			[]byte(".....#...."),
			[]byte("....#...O#"),
			[]byte(".....##..."),
			[]byte("..O#......"),
			[]byte(".....OOO#."),
			[]byte(".O#...O#.#"),
			[]byte("....O#...O"),
			[]byte(".......OOO"),
			[]byte("#..OO###.."),
			[]byte("#.OOO#...O"),
		},
		{
			[]byte(".....#...."),
			[]byte("....#...O#"),
			[]byte(".....##..."),
			[]byte("..O#......"),
			[]byte(".....OOO#."),
			[]byte(".O#...O#.#"),
			[]byte("....O#...O"),
			[]byte(".......OOO"),
			[]byte("#...O###.O"),
			[]byte("#.OOO#...O"),
		},
	}

	for i, want := range afterCycle {
		p = cycle(p)
		if !reflect.DeepEqual(p, want) {
			t.Errorf("cycle #%d", i)
			for r := 0; r < min(len(p), len(want)); r++ {
				t.Errorf("%-4d %s   %s", r, string(p[r]), string(want[r]))
				return
			}
		}
	}
}

func Test_hashPlatform(t *testing.T) {
	ps := []platform{
		{
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
		{
			[]byte(".....#...."),
			[]byte("....#...O#"),
			[]byte("...OO##..."),
			[]byte(".OO#......"),
			[]byte(".....OOO#."),
			[]byte(".O#...O#.#"),
			[]byte("....O#...."),
			[]byte("......OOOO"),
			[]byte("#...O###.."),
			[]byte("#..OO#...."),
		},
		{
			[]byte(".....#...."),
			[]byte("....#...O#"),
			[]byte(".....##..."),
			[]byte("..O#......"),
			[]byte(".....OOO#."),
			[]byte(".O#...O#.#"),
			[]byte("....O#...O"),
			[]byte(".......OOO"),
			[]byte("#..OO###.."),
			[]byte("#.OOO#...O"),
		},
		{
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
	}
	hashes := []uint64{}
	for _, p := range ps {
		hashes = append(hashes, getHash(p))
	}

	if hashes[0] != hashes[3] {
		for i := range hashes {
			t.Errorf("hash[%d]=%d", i, hashes[i])
		}
	}
}

func Test_cycles(t *testing.T) {
	p := platform{
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
	}

	got := cyclesAndWeight(p, 1000000000)
	want := 64

	if got != want {
		t.Errorf("cyclesAndWeight(cycles(1000000000))=%d, want:%d", got, want)
	}
}
