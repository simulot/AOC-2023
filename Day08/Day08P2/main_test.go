package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_getSteps(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want int
	}{
		{
			name: "map 1",
			m: readMap(internal.ReadString(`
			LR

			11A = (11B, XXX)
			11B = (XXX, 11Z)
			11Z = (11B, XXX)
			22A = (22B, XXX)
			22B = (22C, 22C)
			22C = (22Z, 22Z)
			22Z = (22B, 22B)
			XXX = (XXX, XXX)`)),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSteps(tt.m); got != tt.want {
				t.Errorf("getSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{10, 15},
			want:  30,
		},
		{
			input: []int{10, 15, 20},
			want:  60,
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  2520,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := lcm(tt.input...)
			if got != tt.want {
				t.Errorf("lcm()=%v want %v", got, tt.want)
			}
		})
	}
}
