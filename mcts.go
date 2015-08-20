package main

import (
	"fmt"
)

func uct(initialState State, iterations int) int {
	root := &Node{
		PlayerThatMoved: initialState.PlayerThatMoved(),
		RemainingMoves: initialState.PossibleMoves(),
	}
	for i := 0; i < iterations; i++ {
		node := root
		state := initialState.Copy()

		// Select
		for len(node.RemainingMoves) == 0 && len(node.Children) > 0 {
			node = node.SelectMostPromisingNode()
			state.PerformMove(node.Move)
		}

		// Expand
		if len(node.RemainingMoves) > 0 {
			move := node.PickRandomRemainingMove()
			state.PerformMove(move)
			node = node.AddChild(move, state)
		}

		// Rollout
		if len(state.PossibleMoves()) > 1 {
			for state.PerformRandomMove() {}
		}

		// Backpropagate
		for node != nil {
			node.Update(state.Winner(node.PlayerThatMoved))
			node = node.Parent
		}
	}

	//root.Debug()

	// Return the best move
	return root.MostVisitedChild().Move
}

func Play(state State, debug bool) {
	for len(state.PossibleMoves()) > 0 {
		move := uct(state, 50)
		state.PerformMove(move)
		if debug {
			fmt.Printf("Player %v took move %v\n\n", state.PlayerThatMoved(), move)
		}
	}

}

func main() {
	state := nim(15)
	Play(state, true)
	if state.Winner(0) {
		fmt.Println("Player 0 won")
	} else {
		fmt.Println("Player 1 won")
	}
}
