package mcts

import (
	"fmt"
)

type mcts struct {
	K float64
	Debug bool
}

func New() *mcts {
	return &mcts{
		K: 1.0,
		Debug: false,
	}
}

func (m *mcts) Find(initialState State, iterations int) Move {
	root := &node{
		PlayerThatMoved: initialState.PlayerThatMoved(),
		RemainingMoves: initialState.PossibleMoves(),
	}
	for i := 0; i < iterations; i++ {
		node := root
		state := initialState.CopyRandomized()

		// Select
		for len(node.RemainingMoves) == 0 && len(node.Children) > 0 {
			node = node.SelectMostPromisingNode(m.K)
			node.Move.Perform(state)
		}

		// Expand
		if len(node.RemainingMoves) > 0 {
			move := node.PickRandomRemainingMove()
			move.Perform(state)
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

func (m *mcts) Play(state State, iterations int) {
	for len(state.PossibleMoves()) > 0 {
		move := m.Find(state, iterations)
		move.Perform(state)
		if m.Debug {
			fmt.Printf("Player %v took move %v\n\n", state.PlayerThatMoved(), move)
		}
	}

}
