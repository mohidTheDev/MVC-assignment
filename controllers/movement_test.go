package controllers

import (
	"coclone/models"
	"testing"
)

func TestFindDistance(t *testing.T) {
	troop := &models.Troop{X: 0, Y: 0}
	structure := &models.Structure{X: 3, Y: 4}

	dist := findDistance(troop, structure)
	if dist != 5.0 {
		t.Errorf("Expected distance 5.0, got %f", dist)
	}
}

func TestAStarPathfinding(t *testing.T) {
	gridSize := 5
	state := &models.BattleState{
		Grid: make([][]*models.Cell, gridSize),
	}
	for i := 0; i < gridSize; i++ {
		state.Grid[i] = make([]*models.Cell, gridSize)
		for j := 0; j < gridSize; j++ {
			state.Grid[i][j] = &models.Cell{X: i, Y: j, Wall: false}
		}
	}

	path, blockingWall := aStar(0, 0, 0, 2, state)

	if blockingWall != "" {
		t.Errorf("Expected no blocking wall, got '%s'", blockingWall)
	}
	if len(path) == 0 {
		t.Errorf("Expected a valid path, got empty array")
	}

	state.Grid[0][1].Wall = true
	state.Grid[0][1].StructureID = "wall_1"

	path2, blockingWall2 := aStar(0, 0, 0, 2, state)

	if len(path2) == 0 && blockingWall2 == "" {
		t.Errorf("A* failed to route around single wall or identify blockage properly")
	}
}

func TestCalcH(t *testing.T) {
	h := calcH(0, 0, 3, 4)

	expected := 52
	if h != expected {
		t.Errorf("Expected heuristic %d, got %d", expected, h)
	}
}
