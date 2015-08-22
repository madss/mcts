package mcts

import (
	"fmt"
	"math"
	"math/rand"
)

type node struct {
	Parent         *node
	Children       []*node
	Move           Move
	RemainingMoves []Move
	PlayerThatMoved int
	Wins           int
	Visits         int
}

func (n *node) SelectMostPromisingNode() *node {
	var bestChild *node
	var bestScore float64 = 0.0

	for _, child := range n.Children {
		k := 1.0
		cWins := float64(child.Wins)
		cVisits := float64(child.Visits)
		visits := float64(n.Visits)
		score := cWins/cVisits + k*math.Sqrt(2.0*math.Log(visits)/cVisits)
		if score > bestScore {
			bestChild = child
			bestScore = score
		}
	}
	return bestChild
}

func (n *node) PickRandomRemainingMove() Move {
	length := len(n.RemainingMoves)
	index := rand.Intn(length)
	move := n.RemainingMoves[index]

	// Remove the selected move
	n.RemainingMoves[index] = n.RemainingMoves[length-1]
	n.RemainingMoves[length - 1] = nil  // avoid memory leaks
	n.RemainingMoves = n.RemainingMoves[:length-1]

	return move
}

func (n *node) AddChild(move Move, state State) *node {
	newNode := &node{
		Parent:         n,
		Move:           move,
		PlayerThatMoved: state.PlayerThatMoved(),
		RemainingMoves: state.PossibleMoves(),
	}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *node) Update(won bool) {
	if won {
		n.Wins += 1
	}
	n.Visits++
}

func (n *node) MostVisitedChild() *node {
	var mostVisits int
	var mostVisited *node
	for _, child := range n.Children {
		if child.Visits > mostVisits {
			mostVisits = child.Visits
			mostVisited = child
		}
	}
	return mostVisited
}

func (n *node) Debug() {
	for _, child := range n.Children {
		fmt.Printf("Move: %v, wins: %v, visits: %v\n", child.Move, child.Wins, child.Visits)
	}
}
