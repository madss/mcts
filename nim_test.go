package main

import "testing"

func TestNim(t *testing.T) {
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
		state := nim(c.chips)
		Play(state, false)
		if !state.Winner(c.winner) {
			t.Error("Player %d should win with %d chips", c.winner, c.chips)
		}
	}
}
