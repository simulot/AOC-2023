package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_getProcessHistory(t *testing.T) {
	tests := []struct {
		name    string
		history []int
		want    int
	}{
		{
			name:    "line1",
			history: internal.ScanNumbers("0 3 6 9 12 15"),
			want:    18,
		},
		{
			name:    "line2",
			history: internal.ScanNumbers("1 3 6 10 15 21"),
			want:    28,
		},
		{
			name:    "line3",
			history: internal.ScanNumbers("10 13 16 21 30 45"),
			want:    68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processHistory(tt.history); got != tt.want {
				t.Errorf("processHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
