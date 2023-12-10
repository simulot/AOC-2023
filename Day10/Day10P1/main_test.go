package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_readMaze(t *testing.T) {
	tests := []struct {
		name          string
		lines         []string
		wantStart     point
		wantStartTile pipeTile
	}{
		{
			name: "test1",
			lines: internal.ReadString(`
			.....
			.S-7.
			.|.|.
			.L-J.
			.....`),
			wantStart:     point{vertical: 1, horizontal: 1},
			wantStartTile: 'F',
		},
		{
			name: "test2",
			lines: internal.ReadString(`
			-L|F7
			7S-7|
			L|7||
			-L-J|
			L|-JF`),
			wantStart:     point{vertical: 1, horizontal: 1},
			wantStartTile: 'F',
		},
		{
			name: "test3",
			lines: internal.ReadString(`
			7-F7-
			.FJ|7
			SJLL7
			|F--J
			LJ.LJ`),
			wantStart:     point{vertical: 2, horizontal: 0},
			wantStartTile: 'F',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := readMaze(tt.lines)
			if tt.wantStart != m.start {
				t.Errorf("Test_readMaze()  start =%v, want %v", m.start, tt.wantStart)
			}
			if tt.wantStartTile != m.tileAt(m.start) {
				t.Errorf("Test_readMaze()  startTile =%s, want %s", m.tileAt(m.start), tt.wantStartTile)
			}
		})
	}
}

func Test_getFarthestPosition(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "test1",
			lines: internal.ReadString(`
			.....
			.S-7.
			.|.|.
			.L-J.
			.....`),
			want: 4,
		},
		{
			name: "test2",
			lines: internal.ReadString(`
			..F7.
			.FJ|.
			SJ.L7
			|F--J
			LJ...`),
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFarthestPosition(readMaze(tt.lines))
			if tt.want != got {
				t.Errorf("Test_getFarthestPosition()=%v, want %v", got, tt.want)
			}
		})
	}
}
