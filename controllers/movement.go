package controllers

import (
	"coclone/models"
	"math"
)

func findDistance(troop *models.Troop, structure *models.Structure) float64 {
	return math.Hypot(troop.X-structure.X, troop.Y-structure.Y)
}

type pathNode struct {
	cell   *models.Cell
	parent *pathNode
	gCost  int
	hCost  int
}

type coord struct {
	x, y int
}

func aStar(startX, startY, targetX, targetY int, state *models.BattleState) []*models.Cell {
	path := []*models.Cell{}
	//Use a* algorithm to find path
	return path
}
