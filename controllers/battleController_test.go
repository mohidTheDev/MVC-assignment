package controllers

import (
	"coclone/models"
	"testing"
)

func TestIsBattleOver(t *testing.T) {
	state1 := &models.BattleState{
		Structures: map[string]*models.Structure{
			"1": {Name: "Cannon"},
			"2": {Name: "Wall"},
		},
		Troops: map[string]*models.Troop{
			"t_1": {Name: "Barbarian"},
		},
	}
	if isBattleOver(state1) == true {
		t.Errorf("Expected battle to continue, but it ended prematurely.")
	}

	state2 := &models.BattleState{
		Structures: map[string]*models.Structure{
			"1": {Name: "Wall"},
		},
		Troops: map[string]*models.Troop{
			"t_1": {Name: "Barbarian"},
		},
	}
	if isBattleOver(state2) == false {
		t.Errorf("Expected battle to end because only walls remain.")
	}

	state3 := &models.BattleState{
		Structures: map[string]*models.Structure{
			"1": {Name: "Cannon"},
		},
		Troops: make(map[string]*models.Troop),
	}
	if isBattleOver(state3) == false {
		t.Errorf("Expected battle to end because all attacking troops are dead.")
	}
}
