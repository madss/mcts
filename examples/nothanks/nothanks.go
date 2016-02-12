package main

import (
	"fmt"
	"github.com/madss/mcts"
)

type NoThanksState struct {
	players []*player
	drawPile []deck
	currentCard card
	currentChips int
}

func New(noPlayers int) *NoThanksState {
	players := make([]player, noPlayers)
	for i := 0; i < noPlayers; i++ {
		players[i] = &player{ chips = 11 }
	}
	return &NoThanksState{
		drawPile: createDeck(3, 35)
	}
}

func (s *NoThanksState) CopyRandomized() mcts.State {
	return s
}

func (s *NoThanksState) PlayerThatMoved() int {
	return 0
}

func (s *NoThanksState) PossibleMoves() []mcts.Move {
	return []mcts.Move{}
}

func (s *NoThanksState) PerformRandomMove() bool {
	return false
}

func (s *NoThanksState) Winner(player int) bool {
	return true
}

func (s *NoThanksState) String() string {
	return ""
}

func main() {
	state := New(5)
	ts := mcts.New()
	ts.Play(state, 1000)
	for p := 0; p < 5; p++ {
		if state.Winner(p) {
			fmt.Printf("Player %d won\n", p)
		}
	}
}
