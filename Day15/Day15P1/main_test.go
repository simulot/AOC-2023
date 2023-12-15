package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_hash(t *testing.T) {
	tc := []struct {
		name string
		want int
	}{
		{
			name: "HASH",
			want: 52,
		},
		{
			name: "rn=1",
			want: 30,
		},
		{
			name: "cm-",
			want: 253,
		},
		{name: "qp=3", want: 97},
		{name: "cm=2", want: 47},
		{name: "qp-", want: 14},
		{name: "pc=4", want: 180},
		{name: "ot=9", want: 9},
		{name: "ab=5", want: 197},
		{name: "pc-", want: 48},
		{name: "pc=6", want: 214},
		{name: "ot=7", want: 231},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := hash(tt.name)
			if got != tt.want {
				t.Errorf("hash('%s')=%d, want: %d", tt.name, got, tt.want)
			}
		})
	}
}

func Test_sumOfHashes(t *testing.T) {
	tc := []struct {
		name string
		s    []string
		want int
	}{
		{
			name: "test set",
			s:    internal.CSVSplit("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"),
			want: 1320,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfHashes(tt.s)
			if got != tt.want {
				t.Errorf("sumOfHashes()=%d, want: %d", got, tt.want)
			}
		})
	}
}
