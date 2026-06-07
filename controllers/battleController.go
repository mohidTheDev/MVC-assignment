package controllers

import (
	"coclone/models"
	"time"
)

func StartBattle(state *models.BattleState) {
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
