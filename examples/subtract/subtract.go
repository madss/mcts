package main

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/madss/mcts"
)

type SubtractState struct {
	player int
	chips  int
}

func subtract(chips int) mcts.State {
	return &SubtractState{1, chips}
}

func (s *SubtractState) CopyRandomized() mcts.State {
	copy := *s
	return &copy
}

func (s *SubtractState) PlayerThatMoved() int {
	return s.player
}

func (s *SubtractState) PossibleMoves() []mcts.Move {
	count := int(math.Min(4.0, float64(s.chips)))
	moves := make([]mcts.Move, count)
	for i := 0; i < count; i++ {
		moves[i] = SubtractMove{ chips: i + 1 }
	}
	return moves
}

func (s *SubtractState) PerformRandomMove() bool {
	count := int(math.Min(4.0, float64(s.chips)))
	SubtractMove{ chips: rand.Intn(count) + 1 }.Perform(s)
	return s.chips > 0
}

func (s *SubtractState) Winner(player int) bool {
	return player == s.player
}

func (s *SubtractState) Debug() {
	fmt.Printf("Player: %d, Chips: %d\n", s.player, s.chips)
}

type SubtractMove struct {
	chips int
}

func (s SubtractMove) Perform(state mcts.State) {
	subtractState, ok := state.(*SubtractState)
	if ok {
		subtractState.chips -= s.chips
		subtractState.player = (subtractState.player + 1) % 2
	}
}

func (s SubtractMove) String() string {
	return fmt.Sprintf("Take %d chips", s.chips)
}

func main() {
	state := subtract(15)
	mcts.New().Play(state, 50)
	if state.Winner(0) {
		fmt.Println("Player 0 won")
	} else {
		fmt.Println("Player 1 won")
	}
}
