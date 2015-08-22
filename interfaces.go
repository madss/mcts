package mcts

// State defines methods required to perform a monte carlo tree search.
type State interface {
	// CopyRandomized creates a copy of the current state with all
	// unknown state randomized. For games with absolute information
	// randomization doesn't matter, but for a card game, for instance, all
	// unrevealed cards must be shuffled to avoid biasing the suggested move
	// towards a particular sequence of cards.
	CopyRandomized() State

	// PlayerThatMoved returns the zero-based index of the player that
	// performed the most recent move. The value before the first move is
	// performed is irrelevant, but right after the first move is performed it
	// should return zero.
	PlayerThatMoved() int

	// PossibleMoves returns a list of moves currently available.
	PossibleMoves() []Move

	// PerformRandomMove updates the state to reflect a randomly chosen move.
	// The implementation should be equivalent to
	//
	//     moves := state.PossibleMoves()
	//     index := rand.Intn(len(moves))
	//     moves[index].Perform(state)
	//
	// but allows for optimizations by not requiring all moves to be generated.
	// It returns whether more moves are available, equivalent to
	//
	//     len(state.PossibleMoves()) > 0
	PerformRandomMove() bool

	// Winner return whether the nth player is considered a winner based on the
	// current state of the game. The behaviour is until there are no more
	// possible moves, in which case the game is considered to be over.
	Winner(int) bool
}

// Move implements a specific move that can be performed on a given state.
type Move interface {
	// Perform updates the state to reflect the move.
	Perform(State)

	// String return a string representation of the move.
	String() string
}
