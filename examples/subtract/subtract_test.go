package main

import (
	"testing"
	"github.com/madss/mcts"
)

func TestSubtract(t *testing.T) {
	cases := []struct {
		chips, winner int
	}{
		{  1, 0 },
		{  2, 0 },
		{  3, 0 },
		{  4, 0 },
		{  5, 1 },
		{  6, 0 },
		{  7, 0 },
		{  8, 0 },
		{  9, 0 },
		{ 10, 1 },
	}
	for _, c := range cases {
		state := subtract(c.chips)
		mcts.New().Play(state, 50)
		if !state.Winner(c.winner) {
			t.Error("Player %d should win with %d chips", c.winner, c.chips)
		}
	}
}
