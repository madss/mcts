package main

// State defines methods required to perform a monte carlo tree search
type State interface {
	// CopyRandomized should create a copy of the current state with all
	// unknown state randomized. For games with absolute information
	// randomization doesn't matter, but for a card game, for instance, all
	// unrevealed cards must be shuffled to avoid biasing the suggested move
	// towards a particular sequence of cards.
	CopyRandomized() State

	// PlayerThatMoved should return the zero-based index of the player that
	// performed the most recent move. The value before the first move is
	// performed is irrelevant, but after the first move is performed it should
	// return zero.
	PlayerThatMoved() int

	// PossibleMoves should return a list of moves currently available.
	PossibleMoves() []int

	// PerformMove updates the state to reflect the move.
	PerformMove(int)

	// PerformRandomMove updates the state to reflect a randomly chosen move.
	// The implementation should in practice be equivalent to
	//
	//     moves := state.PossibleMoves()
	//     index := rand.Intn(len(moves))
	//     state.PerformMove(moves[index])
	//
	// It returns whether more moves are availble, equivalent to
	//
	//     len(state.PossibleMoves()) > 0
	PerformRandomMove() bool

	// Winner return whether the nth player is considered a winner based on the
	// current state of the game. The behaviour is until there are no more
	// possible moves, in which case the game is considered to be over.
	Winner(int) bool
}
