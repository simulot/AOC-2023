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
			name: "71530-940200",
			race: race{t: 71530, d: 940200},
			want: 71503,
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
