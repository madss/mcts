package main

import (
	"bytes"
	"fmt"
 	"math/rand"
	"mcts"
)

type SubtractState struct {
	player int
	chips  int
}

func subtract(chips int) *SubtractState {
	return &SubtractState{ chips: chips }
}

func (s *SubtractState) Copy() mcts.State {
	copy := *s
	return &copy
}

func (s *SubtractState) CurrentPlayer() int {
	return s.player
}

func (s *SubtractState) PossibleMoveCount() int {
	moves := 4
	if s.chips < moves {
		moves = s.chips
	}
	return moves
}

func (s *SubtractState) PossibleMoves() []mcts.Move {
	count := s.PossibleMoveCount()
	moves := make([]mcts.Move, count)
	for i := 0; i < count; i++ {
		moves[i] = SubtractMove{ chips: i + 1 }
	}
	return moves
}

func (s *SubtractState) PerformRandomMove() bool {
	count := s.PossibleMoveCount()
	if count > 0 {
		SubtractMove{ chips: rand.Intn(count) + 1 }.Perform(s)
		return true
	}
	return false
}

func (s *SubtractState) Winner(player int) bool {
	return player == s.player
}

func (s *SubtractState) String() string {
	var buffer bytes.Buffer
	for i := 0; i < s.chips; i++ {
		buffer.WriteString("*")
	}
	fmt.Fprintf(&buffer, " (%d chips)", s.chips)
	return buffer.String()
}

type SubtractMove struct {
	chips int
}

func (s SubtractMove) Perform(state mcts.State) {
	subtractState := state.(*SubtractState)
	subtractState.chips -= s.chips
	if subtractState.chips > 0 {
		subtractState.player = (subtractState.player + 1) % 2
	}
}

func main() {
	state := subtract(15)
	for state.PossibleMoveCount() > 0 {
	  var move SubtractMove
	  fmt.Println(state.String())
	  if state.CurrentPlayer() > 0 {
			move = mcts.Suggest(state, 50, 1.0).(SubtractMove)
			move.Perform(state)
			fmt.Printf("Computer took %d chips\n", move.chips)
		} else {
			fmt.Print("How many chips do you take: ")
			fmt.Scan(&move.chips)
			move.Perform(state)
		}
	}
	if state.Winner(0) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost")
	}
}
