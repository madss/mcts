package main

import (
	"fmt"
	"math"
	"math/rand"
)

type NimState struct {
	player int
	chips  int
}

func nim(chips int) State {
	return &NimState{1, chips}
}

func (n *NimState) CopyRandomized() State {
	copy := *n
	return &copy
}

func (n *NimState) PlayerThatMoved() int {
	return n.player
}

func (n *NimState) PossibleMoves() []Move {
	count := int(math.Min(4.0, float64(n.chips)))
	moves := make([]Move, count)
	for i := 0; i < count; i++ {
		moves[i] = NimMove{ chips: i + 1 }
	}
	return moves
}

func (n *NimState) PerformRandomMove() bool {
	count := int(math.Min(4.0, float64(n.chips)))
	NimMove{ chips: rand.Intn(count) + 1 }.Perform(n)
	return n.chips > 0
}

func (n *NimState) Winner(player int) bool {
	return player == n.player
}

func (n *NimState) Debug() {
	fmt.Printf("Player: %d, Chips: %d\n", n.player, n.chips)
}

type NimMove struct {
	chips int
}

func (n NimMove) Perform(state State) {
	s, ok := state.(*NimState)
	if ok {
		s.chips -= n.chips
		s.player = (s.player + 1) % 2
	}
}

func (n NimMove) String() string {
	return fmt.Sprintf("Take %d chips", n.chips)
}
