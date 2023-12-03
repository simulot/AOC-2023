package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_validGames(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "sample",
			lines: internal.ReadString(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`),
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validGames(tt.lines, map[color]int{
				red:   12,
				green: 13,
				blue:  14,
			},
			); got != tt.want {
				t.Errorf("validGames() = %v, want %v", got, tt.want)
			}
		})
	}
}
