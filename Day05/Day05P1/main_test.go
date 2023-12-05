package main

import (
	"AOC2023/internal"
	"fmt"
	"testing"
)

func Test_mapSeedToSoil(t *testing.T) {
	m := []Destination{
		{To: 50, From: 98, Length: 2},
		{To: 52, From: 50, Length: 48},
	}
	tests := []struct {
		seed int
		want int
	}{
		{seed: 0, want: 0},
		{seed: 1, want: 1},
		{seed: 48, want: 48},
		{seed: 49, want: 49},
		{seed: 50, want: 52},
		{seed: 51, want: 53},
		{seed: 96, want: 98},
		{seed: 97, want: 99},
		{seed: 98, want: 50},
		{seed: 99, want: 51},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("seed %d to soil %d", tt.seed, tt.want), func(t *testing.T) {
			got := mapSeedToSoil(m, tt.seed)
			if got != tt.want {
				t.Errorf("seedLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nearestSoil(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "sample",
			lines: internal.ReadString(`
			seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48

			soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15

			fertilizer-to-water map:
			49 53 8
			0 11 42
			42 0 7
			57 7 4

			water-to-light map:
			88 18 7
			18 25 70

			light-to-temperature map:
			45 77 23
			81 45 19
			68 64 13

			temperature-to-humidity map:
			0 69 1
			1 0 69

			humidity-to-location map:
			60 56 37
			56 93 4
			`),
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestSoil(readAlmanac(tt.lines)); got != tt.want {
				t.Errorf("nearestSoil() = %v, want %v", got, tt.want)
			}
		})
	}
}
