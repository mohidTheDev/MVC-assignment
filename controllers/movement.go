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

type aNode struct {
	cell    *models.Cell
	parent  *aNode
	g, h, f int
}

func aStar(startX, startY, targetX, targetY int, state *models.BattleState) []*models.Cell {
	grid := state.Grid

	type pos [2]int

	openSet := make(map[pos]*aNode)
	closedSet := make(map[pos]bool)

	startPos := pos{startX, startY}
	startNode := &aNode{
		cell: grid[startX][startY],
		h:    calcH(startX, startY, targetX, targetY),
	}
	startNode.f = startNode.h

	openSet[startPos] = startNode
	closest := startNode

	dirs := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for len(openSet) > 0 {
		var current *aNode
		var currentPos pos
		minF := 1000

		for p, node := range openSet {
			if node.f < minF {
				minF = node.f
				current = node
				currentPos = p
			} else if node.f == minF && current != nil && node.h < current.h {
				current = node
				currentPos = p
			}
		}

		delete(openSet, currentPos)
		closedSet[currentPos] = true

		if currentPos[0] == targetX && currentPos[1] == targetY {
			return buildPath(current)
		}

		if current.h < closest.h {
			closest = current
		}

		for _, d := range dirs {
			nx := currentPos[0] + d[0]
			ny := currentPos[1] + d[1]
			np := pos{nx, ny}

			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[nx]) || grid[nx][ny].Wall || closedSet[np] {
				continue
			}

			cost := 10
			if d[0] != 0 && d[1] != 0 {
				cost = 14
			}
			tentativeG := current.g + cost

			nNode, exists := openSet[np]
			if !exists || tentativeG < nNode.g {
				if !exists {
					nNode = &aNode{cell: grid[nx][ny]}
					openSet[np] = nNode
				}
				nNode.parent = current
				nNode.g = tentativeG
				nNode.h = calcH(nx, ny, targetX, targetY)
				nNode.f = nNode.g + nNode.h
			}
		}
	}

	return buildPath(closest)
}

func calcH(x1, y1, x2, y2 int) int {
	dx, dy := x1-x2, y1-y2
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx > dy {
		return 14*dy + 10*(dx-dy)
	}
	return 14*dx + 10*(dy-dx)
}

func buildPath(node *aNode) []*models.Cell {
	var path []*models.Cell

	for curr := node; curr != nil; curr = curr.parent {
		path = append(path, curr.cell)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
