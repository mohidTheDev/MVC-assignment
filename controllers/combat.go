package controllers

import (
	"coclone/models"
	"math"
)

func updateTroop(troop *models.Troop, state *models.BattleState) {
	if troop.TargetID == "" {
		troop.TargetID = findNearestTarget(troop.X, troop.Y, true, state)
	}
	targetDistance := findDistance(troop, state.Structures[troop.TargetID])
	attackRange := 1.0
	if !troop.IsMelee {
		attackRange = troop.AttackRange
	}
	if targetDistance <= attackRange {
		//Attack the target
		return
	}
	//find path to target using a*
	//move <movespeed> steps to target
}

func updateStructure(structure *models.Structures, state *models.BattleState) {
	if structure.TargetID == "" {
		structure.TargetID = findNearestTarget(structure.X, structure.Y, false, state)
	}
	targetDistance := findDistance(state.Troops[structure.TargetID], structure)
	if targetDistance <= structure.AttackRange {
		//Attack the target
	}
}

func findDistance(troop *models.Troop, structure *models.Structures) float64 {
	return math.Hypot(troop.X-structure.X, troop.Y-structure.Y)
}

func findNearestTarget(posX, posY float64, isAttacker bool, state *models.BattleState) string {
	if isAttacker {
		minDistance := 1000.0
		var closestTargetID string
		for _, structure := range state.Structures {
			distance := math.Hypot(posX-structure.X, posY-structure.Y)
			if distance < minDistance {
				minDistance = distance
				closestTargetID = structure.ID
			}
		}
		return closestTargetID
	}
	minDistance := 1000.0
	var closestTargetID string
	for _, troop := range state.Troops {
		distance := math.Hypot(posX-troop.X, posY-troop.Y)
		if distance < minDistance {
			minDistance = distance
			closestTargetID = troop.ID
		}
	}
	return closestTargetID
}
