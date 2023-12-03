package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_trebuchet(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "sample",
			lines: internal.ReadString(`two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`),
			want: 281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trebuchet(tt.lines); got != tt.want {
				t.Errorf("trebuchet() = %v, want %v", got, tt.want)
			}
		})
	}
}
