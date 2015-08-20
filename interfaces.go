package main

type State interface {
	Copy() State
	PlayerThatMoved() int
	PossibleMoves() []int
	PerformMove(int)
	PerformRandomMove() bool
	Winner(int) bool
	Debug()
}
