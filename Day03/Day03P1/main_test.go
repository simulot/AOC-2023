package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_parNumbers(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "sample",
			lines: internal.ReadString(`467..114..
			...*......
			..35..633.
			......#...
			617*......
			.....+.58.
			..592.....
			......755.
			...$.*....
			.664.598..
			`),
			want: 4361,
		},

		{
			name: "sample3",
			lines: internal.ReadString(`12.......*..
			+.........34
			.......-12..
			..78........
			..*....60...
			78..........
			.......23...
			....90*12...
			............
			2.2......12.
			.*.........*
			1.1.......56`),
			want: 413,
		},
		{
			name: "sample4",
			lines: internal.ReadString(`12.......*..
			+.........34
			.......-12..
			..78........
			..*....60...
			78.........9
			.5.....23..$
			8...90*12...
			............
			2.2......12.
			.*.........*
			1.1..503+.56`),
			want: 925,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parNumbers(tt.lines); got != tt.want {
				t.Errorf("parNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
