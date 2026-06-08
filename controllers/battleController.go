package controllers

import (
	"coclone/models"
	"time"
)

func StartBattle(state *models.BattleState) {
	//Fetch village layout and troops from DB

	go runBattleLoop(state)
}

func runBattleLoop(state *models.BattleState) {
	ticker := time.NewTicker(100 * time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			update(state)

			if isBattleOver(state) {
				return
			}
		}
	}
}

func update(state *models.BattleState) {
	for _, troop := range state.Troops {
		updateTroop(troop, state)
	}
	for _, structure := range state.Structures {
		updateStructure(structure, state)
	}
}

func isBattleOver(state *models.BattleState) bool {
	if len(state.Structures) == 0 || len(state.Troops) == 0 {
		return true
	}
	return false
}

func initiateVillage(state *models.BattleState) {
	for _, structure := range state.Structures {
		if structure.Name == "Wall" {
			addWallToGrid(structure, state)
		}
	}
	// Get village buildings from database

	// Dummy data for testing
	state.Structures["1"] = &models.Structure{
		ID:    "1",
		Name:  "Wall",
		X:     5,
		Y:     5,
		HP:    100,
		MaxHP: 100,
	}
	state.Structures["2"] = &models.Structure{
		ID:          "2",
		Name:        "Cannon",
		X:           7,
		Y:           7,
		Damage:      20,
		AttackRange: 3,
		HP:          150,
		MaxHP:       150,
	}
	state.Structures["3"] = &models.Structure{
		ID:          "3",
		Name:        "Cannon",
		X:           9,
		Y:           9,
		Damage:      20,
		AttackRange: 3,
		HP:          150,
		MaxHP:       150,
	}

	//For the sake of testing, we will add some troops here. In real implementation, troops will be added when player deploys them
	state.Troops["1"] = &models.Troop{
		ID:          "1",
		Name:        "Barbarian",
		X:           0,
		Y:           0,
		HP:          50,
		MaxHP:       50,
		Damage:      10,
		MoveSpeed:   1,
		IsMelee:     true,
		AttackRange: 1,
	}
	state.Troops["2"] = &models.Troop{
		ID:          "2",
		Name:        "Archer",
		X:           0,
		Y:           1,
		HP:          30,
		MaxHP:       30,
		Damage:      15,
		MoveSpeed:   1.5,
		IsMelee:     false,
		AttackRange: 3,
	}
	state.Troops["3"] = &models.Troop{
		ID:          "3",
		Name:        "Barbarian",
		X:           0,
		Y:           2,
		HP:          50,
		MaxHP:       50,
		Damage:      10,
		MoveSpeed:   1,
		IsMelee:     true,
		AttackRange: 1,
	}
}

func addWallToGrid(wall *models.Structure, state *models.BattleState) {
	state.Grid[int(wall.X)][int(wall.Y)].Wall = true
}

func removeWallFromGrid(wall *models.Structure, state *models.BattleState) {
	state.Grid[int(wall.X)][int(wall.Y)].Wall = false
}
