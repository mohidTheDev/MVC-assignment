package controllers

import (
	"coclone/models"
	"fmt"
	"math"
)

func updateTroop(troop *models.Troop, state *models.BattleState) {
	if troop.TargetID == "" {
		blockingWallID := ""
		troop.TargetID = findNearestTarget(troop.X, troop.Y, true, state)

		if troop.TargetID == "" {
			return
		}

		troop.MovePath, blockingWallID = aStar(int(troop.X), int(troop.Y), int(state.Structures[troop.TargetID].X), int(state.Structures[troop.TargetID].Y), state)
		if blockingWallID != "" {
			troop.TargetID = blockingWallID
		}
	}
	targetStructure := state.Structures[troop.TargetID]
	if targetStructure == nil {
		troop.TargetID = ""
		return
	}
	targetDistance := findDistance(troop, state.Structures[troop.TargetID])
	attackRange := 1.0
	if !troop.IsMelee {
		attackRange = troop.AttackRange
	}
	if targetDistance <= attackRange {
		damageStructure(state.Structures[troop.TargetID], troop, troop.Damage, state)
		return
	}

	if len(troop.MovePath) > 0 {
		nextStep := troop.MovePath[0]
		targetX := float64(nextStep.X)
		targetY := float64(nextStep.Y)

		dx := targetX - troop.X
		dy := targetY - troop.Y
		distance := math.Hypot(dx, dy)

		dx /= distance
		dy /= distance

		if distance < troop.MoveSpeed {
			troop.MovePath = troop.MovePath[1:]
			fmt.Printf("%s %s (%d) moved to X:%.1f Y:%.1f\n", troop.Name, troop.ID, troop.HP, troop.X, troop.Y)
			return
		}
		troop.X += dx * troop.MoveSpeed
		troop.Y += dy * troop.MoveSpeed
	}
}

func updateStructure(structure *models.Structure, state *models.BattleState) {
	structure.TargetID = findNearestTarget(structure.X, structure.Y, false, state)
	targetTroop := state.Troops[structure.TargetID]
	if targetTroop == nil {
		structure.TargetID = ""
		return
	}
	targetDistance := findDistance(state.Troops[structure.TargetID], structure)
	if targetDistance <= structure.AttackRange {
		attackTroop(state.Troops[structure.TargetID], structure, structure.Damage, state)
	}
}

func damageStructure(structure *models.Structure, source *models.Troop, damage int, state *models.BattleState) {
	structure.HP -= damage
	if structure.HP <= 0 {
		destroyStructure(structure, source, state)
	}
}

func attackTroop(troop *models.Troop, source *models.Structure, damage int, state *models.BattleState) {
	troop.HP -= damage
	if troop.HP <= 0 {
		//Add entry in battlelog
		fmt.Printf("%s %s destroyed by %s %s\n", troop.Name, troop.ID, source.Name, source.ID)
		source.TargetID = ""
		delete(state.Troops, troop.ID)
	}
}

func destroyStructure(structure *models.Structure, source *models.Troop, state *models.BattleState) {
	//Add entry in battlelog
	fmt.Printf("%s %s destroyed by %s %s %d\n", structure.Name, structure.ID, source.Name, source.ID, source.HP)
	if structure.Name == "Wall" {
		removeWallFromGrid(structure, state)
	}
	for _, troop := range state.Troops {
		if troop.TargetID == structure.ID {
			troop.TargetID = ""
		}
	}
	delete(state.Structures, structure.ID)
}

func findNearestTarget(posX, posY float64, isAttacker bool, state *models.BattleState) string {
	if isAttacker {
		minDistance := 1000.0
		var closestTargetID string
		for _, structure := range state.Structures {
			if structure.Name == "Wall" {
				continue
			}
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
