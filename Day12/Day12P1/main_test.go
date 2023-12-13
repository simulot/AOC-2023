package main

import (
	"AOC2023/internal"
	"reflect"
	"testing"
)

func Test_getDamagedGroups(t *testing.T) {
	tc := []struct {
		name string
		want []int
	}{
		{
			name: "#.#.###",
			want: []int{1, 1, 3},
		},
		{
			name: ".#...#....###.",
			want: []int{1, 1, 3},
		},
		{
			name: ".#.###.#.######",
			want: []int{1, 3, 1, 6},
		},
		{
			name: "####.#...#...",
			want: []int{4, 1, 1},
		},
		{
			name: "#....######..#####.",
			want: []int{1, 6, 5},
		},
		{
			name: ".###.##....#",
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := getDamagedGroups(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDamagedGroups=(%v), want: %v", got, tt.want)

			}
		})
	}
}

func Test_getArragements(t *testing.T) {
	tc := []struct {
		report report
		want   []string
	}{
		{
			report: report{
				status:  "?###????????",
				damaged: []int{3, 2, 1},
			},
			want: []string{
				".###.##.#...",
				".###.##..#..",
				".###.##...#.",
				".###.##....#",
				".###..##.#..",
				".###..##..#.",
				".###..##...#",
				".###...##.#.",
				".###...##..#",
				".###....##.#",
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.report.status, func(t *testing.T) {
			got := getArrangements(tt.report)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArrangements()=%v, want:%v", len(got), len(tt.want))
				for i := 0; i < min(len(got), len(tt.want)); i++ {
					t.Logf("%d: %v   %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func Test_getSumOfDamageReports(t *testing.T) {
	reports := readReports(internal.ReadString(`
	???.### 1,1,3
	.??..??...?##. 1,1,3
	?#?#?#?#?#?#?#? 1,3,1,6
	????.#...#... 4,1,1
	????.######..#####. 1,6,5
	?###???????? 3,2,1
	`))
	got := getSumOfDamageReports(reports)
	want := 21
	if got != want {
		t.Errorf("getSumOfDamageReports()=%d, want:%d", got, want)
	}
}
