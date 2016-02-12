package mcts

// State defines methods required to perform a monte carlo tree search.
type State interface {
	// Copy creates a copy of the current state. For games with hidden
	// information, all unrevealed state must be randomized to avoid
	// biasing the suggested move towards, for instance, a particular
	// ordering of cards
	Copy() State

	// CurrentPlayer returns the zero-based index of the player that
	// is about to perform on of the possible moves.
	CurrentPlayer() int

	// PossibleMoves returns a list of moves currently available.
	PossibleMoves() []Move

	// PerformRandomMove updates the state to reflect a randomly chosen move.
	// This allows for optimizations by not requiring all moves to be generated.
	// It returns whether it was able to perform a move or no more moves were
	// available. The implementation should be equivalent to
	//
	//     moves := state.PossibleMoves()
	//     if len(moves) > 0 {
	//         index := rand.Intn(len(moves))
	//         moves[index].Perform(state)
	//         return true
	//     }
	//     return false
	PerformRandomMove() bool

	// Winner return whether the nth player is considered a winner based on the
	// current state of the game. The behaviour is undefined until there are no
	// more possible moves, in which case the game is considered to be over.
	Winner(int) bool
}

// Move implements a specific move that can be performed on a given state.
type Move interface {
	// Perform updates the state to reflect the move.
	Perform(State)
}
