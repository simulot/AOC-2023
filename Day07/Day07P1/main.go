package main

import (
	"AOC2023/internal"
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total winnings: ", getWinnings(readHands(lines)))
}

type kind int

const (
	HighCard kind = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

const cardsStrength = "23456789TJQKA"

type game []hand
type hand struct {
	cards         string
	cardsStrength []byte
	kind          kind
	bid           int
}

func readHands(lines []string) game {
	hands := game{}
	for _, l := range lines {
		if len(l) > 0 {
			hands = append(hands, makeHand(l))
		}
	}
	return hands
}

func makeHand(line string) hand {
	s := strings.Fields(line)
	cards := s[0]
	b, _ := strconv.Atoi(s[1])
	strengths := make([]byte, len(cards))
	for i := range cards {
		strengths[i] = byte(bytes.Index([]byte(cardsStrength), []byte(cards[i:i+1])))
	}
	return hand{cards: cards, bid: b, kind: getHandKind(cards), cardsStrength: strengths}
}

func getHandKind(cards string) kind {
	countMap := map[rune]int{}
	for _, c := range cards {
		countMap[c] = countMap[c] + 1

	}
	counts := make([]int, len(countMap))
	i := 0
	for _, c := range countMap {
		counts[i] = c
		i++
	}
	sort.Slice(sort.IntSlice(counts), func(i, j int) bool {
		return !sort.IntSlice(counts).Less(i, j)
	})
	switch {
	case counts[0] == 5:
		return FiveOfKind
	case counts[0] == 4 && counts[1] == 1:
		return FourOfKind
	case counts[0] == 3 && counts[1] == 2:
		return FullHouse
	case counts[0] == 3 && counts[1] == 1:
		return ThreeOfKind
	case counts[0] == 2 && counts[1] == 2:
		return TwoPair
	case counts[0] == 2 && counts[1] == 1:
		return OnePair
	}
	return HighCard
}

func (g game) Less(i, j int) bool {
	switch {
	case g[i].kind < g[j].kind:
		return true
	case g[i].kind > g[j].kind:
		return false
	}
	return bytes.Compare(g[i].cardsStrength, g[j].cardsStrength) < 0
}
func (g game) Len() int { return len(g) }
func (g game) Swap(i, j int) {
	g[j], g[i] = g[i], g[j]
}

func getWinnings(g game) int {
	sort.Sort(g)
	winnings := 0
	for i := range g {
		winnings += (i + 1) * g[i].bid
	}
	return winnings
}
