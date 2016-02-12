package mcts

import "testing"

func TestSelectMostPromisingNode(t *testing.T) {
	// Todo
}

func PickRandomRemainingMove(t *testing.T) {
	// Todo
}

func TestAddChild(t *testing.T) {
	// Todo
}

func TestUpdate(t *testing.T) {
	n := node{Wins: 10, Visits: 20}

	n.Update(true)
	if n.Wins != 11 || n.Visits != 21 {
		t.Error("Failed to update a win")
	}

	n.Update(false)
	if n.Wins != 11 || n.Visits != 22 {
		t.Error("Failed to update a loss")
	}
}

func TestMostVisitedChild(t *testing.T) {
	var n node

	n = node{
		Children: []*node{},
	}
	if n.MostVisitedChild() != nil {
		t.Error("Failed to return nil for empty list")
	}
	n = node{
		Children: []*node{
			&node{Visits: 42},
			&node{Visits: 100},
			&node{Visits: 3},
		},
	}
	if n.MostVisitedChild() != n.Children[1] {
		t.Error("Failed to find most visited node")
	}
}
