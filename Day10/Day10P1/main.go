package main

import (
	"AOC2023/internal"
	"fmt"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Longest distance: ", getFarthestPosition(readMaze(lines)))
}

type pipeTile byte

func (t pipeTile) String() string {
	return string(byte(t))
}

const (
	pipeNorthSouth_V pipeTile = '|'
	pipeEastWest_H   pipeTile = '-'
	pipeNorthEast_L  pipeTile = 'L'
	pipeNorthWest_J  pipeTile = 'J'
	pipeSouthWest_7  pipeTile = '7'
	pipeSouthEast_F  pipeTile = 'F'
	pipeGround       pipeTile = '.'
	pipeStart        pipeTile = 'S'
	pipeInvalid      pipeTile = 'x'
)

type point struct {
	horizontal, vertical int
}

func (p1 point) add(p2 point) point {
	return point{horizontal: p1.horizontal + p2.horizontal, vertical: p1.vertical + p2.vertical}
}
func (p1 point) sub(p2 point) point {
	return point{horizontal: p1.horizontal - p2.horizontal, vertical: p1.vertical - p2.vertical}
}

var (
	mvtNorth = point{0, -1}
	mvtSouth = point{0, 1}
	mvtEast  = point{1, 0}
	mvtWest  = point{-1, 0}
)

var tileConnections = map[pipeTile][2]point{
	pipeNorthSouth_V: {mvtNorth, mvtSouth},
	pipeEastWest_H:   {mvtEast, mvtWest},
	pipeNorthEast_L:  {mvtNorth, mvtEast},
	pipeSouthWest_7:  {mvtSouth, mvtWest},
	pipeSouthEast_F:  {mvtSouth, mvtEast},
	pipeNorthWest_J:  {mvtNorth, mvtWest},
}

type Maze struct {
	size      int
	maze      [][]pipeTile
	distances [][]int
	start     point
}

func (m Maze) tileAt(at point) pipeTile {
	if at.horizontal >= 0 && at.horizontal < m.size && at.vertical >= 0 && at.vertical < m.size {
		return m.maze[at.vertical][at.horizontal]
	}
	return pipeInvalid
}

func (m Maze) getDistanceAt(at point) int {
	return m.distances[at.vertical][at.horizontal]
}

func (m Maze) setDistanceAt(at point, d int) {
	m.distances[at.vertical][at.horizontal] = d
}

func (m Maze) getConnections(at point) []point {
	connections := []point{}

	for _, move := range tileConnections[m.tileAt(at)] {
		p := at.add(move)
		if m.tileAt(p) != pipeInvalid {
			connections = append(connections, p)
		}
	}
	return connections
}

func readMaze(lines []string) Maze {
	lines = internal.FilterOutEmptyLines(lines)

	size := len(lines[0])
	m := Maze{
		size:      size,
		maze:      make([][]pipeTile, size),
		distances: make([][]int, size),
	}

	for vertical, l := range lines {
		if l != "" {
			r := make([]pipeTile, size)
			for horizontal, tile := range l {
				r[horizontal] = pipeTile(tile)
				if tile == 'S' {
					m.start = point{horizontal: horizontal, vertical: vertical}
				}
			}
			m.maze[vertical] = r
			initDist := make([]int, size)
			for i := 0; i < size; i++ {
				initDist[i] = -1
			}
			m.distances[vertical] = initDist
		}
	}

	// what was the kind of pipe at start position?
	for _, startTile := range []pipeTile{pipeEastWest_H, pipeNorthEast_L, pipeNorthSouth_V, pipeNorthWest_J, pipeSouthEast_F, pipeSouthWest_7} {
		// pretend the start tile is startTile
		m.maze[m.start.vertical][m.start.horizontal] = startTile
		// for this startTile, get possible connections
		connectedPoints := m.getConnections(m.start)
		if len(connectedPoints) != 2 {
			// would be connected out of the grid
			continue
		}

		// check the validity of the startTile by checking  by checking if the 2 connected tiles are connected back to the start point
		countValid := 0
		for _, p1 := range connectedPoints {
			t := m.tileAt(p1)
			if t == pipeGround {
				break
			}
			for _, move := range tileConnections[t] {
				p2 := p1.add(move)
				if p2 == m.start {
					countValid++
				}
			}
		}
		if countValid == 2 {
			break
		}
	}
	return m
}

func getFarthestPosition(m Maze) int {
	tracks := []point{m.start}
	farthest := 0
	m.setDistanceAt(m.start, 0)
	for t := 0; t < len(tracks); t++ {
		p1 := tracks[t]
		d := m.getDistanceAt(p1)
		for _, p2 := range m.getConnections(p1) {
			if m.getDistanceAt(p2) == -1 {
				m.setDistanceAt(p2, d+1)
				farthest = max(farthest, d+1)
				tracks = append(tracks, p2)
			}
		}
	}
	return farthest
}
