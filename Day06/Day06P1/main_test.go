package main

import (
	"testing"
)

func Test_checkRace(t *testing.T) {
	tests := []struct {
		name string
		race race
		want int
	}{
		{
			name: "7-9",
			race: race{t: 7, d: 9},
			want: 4,
		},
		{
			name: "15-40",
			race: race{t: 15, d: 40},
			want: 8,
		},
		{
			name: "30-200",
			race: race{t: 30, d: 200},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRace(tt.race); got != tt.want {
				t.Errorf("checkRace() = %v, want %v", got, tt.want)
			}
		})
	}
}
