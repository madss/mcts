package mcts

func Suggest(initialState State, iterations int, k float64) Move {
	root := &node{
		CurrentPlayer: initialState.CurrentPlayer(),
		RemainingMoves:  initialState.PossibleMoves(),
	}

	for i := 0; i < iterations; i++ {
		node := root
		state := initialState.Copy()

		// Select
		for len(node.RemainingMoves) == 0 && len(node.Children) > 0 {
			node = node.SelectMostPromisingNode(k)
			node.Move.Perform(state)
		}

		// Expand
		if len(node.RemainingMoves) > 0 {
			move := node.PickRandomRemainingMove()
			move.Perform(state)
			node = node.AddChild(move, state)
		}

		// Rollout
		for state.PerformRandomMove() {
		}

		// Backpropagate
		for node.Parent != nil {
			node.Update(state.Winner(node.Parent.CurrentPlayer))
			node = node.Parent
		}
		node.Update(false)
	}

	// Return the best move
	mostVisitedChild := root.MostVisitedChild()
	if mostVisitedChild != nil {
		return mostVisitedChild.Move
	} else {
		return nil
	}
}

func PlayOut(state State, iterations int, k float64) (moves []Move) {
	for {
		move := Suggest(state, iterations, k)
		if move != nil {
			move.Perform(state)
			moves = append(moves, move)
		} else {
			return
		}
	}
}
