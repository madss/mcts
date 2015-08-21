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

func (n *NimState) PossibleMoves() []int {
	count := int(math.Min(4.0, float64(n.chips)))
	moves := make([]int, count)
	for i := 0; i < count; i++ {
		moves[i] = i + 1
	}
	return moves
}

func (n *NimState) PerformMove(m int) {
	n.player = (n.player + 1) % 2
	n.chips -= m
}

func (n *NimState) PerformRandomMove() bool {
	moves := n.PossibleMoves()
	index := rand.Intn(len(moves))
	n.PerformMove(moves[index])
	return n.chips > 0
}

func (n *NimState) Winner(player int) bool {
	return player == n.player
}

func (n *NimState) Debug() {
	fmt.Printf("Player: %d, Chips: %d\n", n.player, n.chips)
}
