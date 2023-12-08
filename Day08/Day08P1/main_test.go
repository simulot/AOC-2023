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
			RL

			AAA = (BBB, CCC)
			BBB = (DDD, EEE)
			CCC = (ZZZ, GGG)
			DDD = (DDD, DDD)
			EEE = (EEE, EEE)
			GGG = (GGG, GGG)
			ZZZ = (ZZZ, ZZZ)`)),
			want: 2,
		},
		{
			name: "map 2",
			m: readMap(internal.ReadString(`
			LLR

			AAA = (BBB, BBB)
			BBB = (AAA, ZZZ)
			ZZZ = (ZZZ, ZZZ)`)),
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
