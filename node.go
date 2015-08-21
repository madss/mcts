package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Node struct {
	Parent         *Node
	Children       []*Node
	Move           Move
	RemainingMoves []Move
	PlayerThatMoved int
	Wins           int
	Visits         int
}

func (n *Node) SelectMostPromisingNode() *Node {
	var bestChild *Node
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

func (n *Node) PickRandomRemainingMove() Move {
	length := len(n.RemainingMoves)
	index := rand.Intn(length)
	move := n.RemainingMoves[index]

	// Remove the selected move
	n.RemainingMoves[index] = n.RemainingMoves[length-1]
	n.RemainingMoves[length - 1] = nil  // avoid memory leaks
	n.RemainingMoves = n.RemainingMoves[:length-1]

	return move
}

func (n *Node) AddChild(move Move, state State) *Node {
	node := &Node{
		Parent:         n,
		Move:           move,
		PlayerThatMoved: state.PlayerThatMoved(),
		RemainingMoves: state.PossibleMoves(),
	}
	n.Children = append(n.Children, node)
	return node
}

func (n *Node) Update(won bool) {
	if won {
		n.Wins += 1
	}
	n.Visits++
}

func (n *Node) MostVisitedChild() *Node {
	var mostVisits int
	var mostVisited *Node
	for _, node := range n.Children {
		if node.Visits > mostVisits {
			mostVisits = node.Visits
			mostVisited = node
		}
	}
	return mostVisited
}

func (n *Node) Debug() {
	for _, child := range n.Children {
		fmt.Printf("Move: %v, wins: %v, visits: %v\n", child.Move, child.Wins, child.Visits)
	}
}
