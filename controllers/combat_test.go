package controllers

import (
	"coclone/models"
	"testing"
)

func TestFindNearestTarget(t *testing.T) {
	state := &models.BattleState{
		Structures: map[string]*models.Structure{
			"s_far":   {ID: "s_far", Name: "Cannon", X: 10, Y: 10},
			"s_close": {ID: "s_close", Name: "TownHall", X: 2, Y: 2},
			"s_wall":  {ID: "s_wall", Name: "Wall", X: 1, Y: 1},
		},
	}

	targetID := findNearestTarget(0, 0, true, state)

	if targetID != "s_close" {
		t.Errorf("Expected target 's_close', got '%s'", targetID)
	}
}

func TestAttackTroopAndDestroy(t *testing.T) {
	state := &models.BattleState{
		Troops: map[string]*models.Troop{
			"t_1": {ID: "t_1", Name: "Barbarian", HP: 15},
		},
		Structures: map[string]*models.Structure{
			"s_1": {ID: "s_1", Name: "Cannon", TargetID: "t_1"},
		},
	}

	troop := state.Troops["t_1"]
	cannon := state.Structures["s_1"]

	attackTroop(troop, cannon, 10, state)
	if state.Troops["t_1"].HP != 5 {
		t.Errorf("Expected troop HP to be 5, got %d", state.Troops["t_1"].HP)
	}

	attackTroop(troop, cannon, 10, state)

	if _, exists := state.Troops["t_1"]; exists {
		t.Errorf("Expected troop to be deleted from state.Troops map")
	}
	if cannon.TargetID != "" {
		t.Errorf("Expected cannon to drop its target, but still targeting '%s'", cannon.TargetID)
	}
}
