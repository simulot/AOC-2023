package main

import (
	"AOC2023/internal"
	"testing"
)

func Test_checkHandType(t *testing.T) {
	tests := []struct {
		name string
		hand string
		want kind
	}{
		{
			name: "Five of a kind",
			hand: "AAAAA",
			want: FiveOfKind,
		},
		{
			name: "Four of a kind",
			hand: "AA8AA",
			want: FourOfKind,
		},
		{
			name: "Full house",
			hand: "23332",
			want: FullHouse,
		},
		{
			name: "Three of a kind",
			hand: "TTT98",
			want: ThreeOfKind,
		},
		{
			name: "Two pair",
			hand: "23432",
			want: TwoPair,
		},
		{
			name: "One pair",
			hand: "A23A4",
			want: OnePair,
		},
		{
			name: "High card",
			hand: "23456",
			want: HighCard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getHandKind(tt.hand)
			if got != tt.want {
				t.Errorf("getHandKind()=%v, want %v", got, tt.want)
			}
		})

	}
}

func Test_checkWinnings(t *testing.T) {
	tests := []struct {
		name string
		game game
		want int
	}{
		{
			name: "game",
			game: readHands(internal.ReadString(`
			32T3K 765
			T55J5 684
			KK677 28
			KTJJT 220
			QQQJA 483`)),
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinnings(tt.game); got != tt.want {
				t.Errorf("getWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
