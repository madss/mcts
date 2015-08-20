package main

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
	node := Node{ Wins: 10, Visits: 20 }

	node.Update(true)
	if node.Wins != 11 || node.Visits != 21 {
		t.Error("Failed to update a win")
	}

	node.Update(false)
	if node.Wins != 11 || node.Visits != 22 {
		t.Error("Failed to update a loss")
	}
}

func TestMostVisitedChild(t *testing.T) {
	node := Node{
		Children: []*Node{
			&Node{ Visits: 42 },
			&Node{ Visits: 100 },
			&Node{ Visits: 3 },
		},
	}
	if node.MostVisitedChild() != node.Children[1] {
		t.Error("Failed to find most visited node")
	}
}

