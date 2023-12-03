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
			lines: internal.ReadString(`1abc2
			pqr3stu8vwx
			a1b2c3d4e5f 
			treb7uchet`),
			want: 142,
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
