package main

import (
	"AOC2023/internal"
	"fmt"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(" energizedCount = ", readContraption(internal.FilterOutEmptyLines(lines)).lightIt().energizedCount())
}

type point struct {
	horizontal, vertical int
}

type beam struct {
	position  point
	direction byte
}

func (b beam) move(direction byte) beam {
	switch direction {
	case '^':
		b.position.vertical--
	case 'v':
		b.position.vertical++
	case '<':
		b.position.horizontal--
	case '>':
		b.position.horizontal++
	}
	b.direction = direction
	return b
}

type tile struct {
	t         byte
	energized bool
	visited   string // beam direction when visited.
}

func (t *tile) out(b beam) []beam {
	if strings.Contains(t.visited, string(b.direction)) {
		// this tile has been already crossed in that direction.
		// we can ignore this beam
		return nil
	}
	t.energized = true
	t.visited += string(b.direction)
	switch t.t {
	case '.':
		return []beam{b.move(b.direction)}
	case '-':
		switch b.direction {
		case '>', '<':
			b = b.move(b.direction)
			return []beam{b}
		case 'v', '^':
			return []beam{
				b.move('<'),
				b.move('>'),
			}
		}
	case '|':
		switch b.direction {
		case '^', 'v':
			b = b.move(b.direction)
			return []beam{b}
		case '<', '>':
			return []beam{
				b.move('^'),
				b.move('v'),
			}
		}
	case '/':
		switch b.direction {
		case '>':
			return []beam{b.move('^')}
		case '<':
			return []beam{b.move('v')}
		case 'v':
			return []beam{b.move('<')}
		case '^':
			return []beam{b.move('>')}
		}
	case '\\':
		switch b.direction {
		case '>':
			return []beam{b.move('v')}
		case '<':
			return []beam{b.move('^')}
		case 'v':
			return []beam{b.move('>')}
		case '^':
			return []beam{b.move('<')}
		}
	}
	return nil
}

type contraption struct {
	tiles [][]tile
}

func readContraption(lines []string) *contraption {
	c := contraption{}
	for _, l := range lines {
		row := make([]tile, len(l))
		for i, t := range l {
			row[i] = tile{t: byte(t)}
		}
		c.tiles = append(c.tiles, row)
	}
	return &c
}

func (c contraption) tileAt(p point) *tile {
	if p.horizontal < 0 || p.horizontal >= len(c.tiles[0]) || p.vertical < 0 || p.vertical >= len(c.tiles) {
		return nil
	}
	return &c.tiles[p.vertical][p.horizontal]
}

func (c *contraption) lightIt() *contraption {
	// round := 0
	paths := []beam{{position: point{horizontal: 0, vertical: 0}, direction: '>'}}
	for len(paths) > 0 {
		b := paths[0]
		paths = paths[1:]
		t := c.tileAt(b.position)
		if t == nil {
			continue
		}
		paths = append(paths, t.out(b)...)
		// round++
		// c.display(round)
	}
	return c
}

func (c contraption) energizedCount() int {
	sum := 0
	for _, r := range c.tiles {
		for _, t := range r {
			if t.energized {
				sum++
			}
		}
	}
	return sum
}

func (c contraption) visualize() []string {
	v := []string{}
	for _, r := range c.tiles {
		b := strings.Builder{}
		for _, t := range r {
			switch t.energized {
			case true:
				b.WriteByte('#')
			default:
				b.WriteByte('.')
			}
		}
		v = append(v, b.String())
	}
	return v
}

func (c contraption) display(round int) {
	fmt.Println("round ", round)
	v := c.visualize()
	for _, l := range v {
		fmt.Println(l)
	}
}
