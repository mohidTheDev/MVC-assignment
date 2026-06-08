package controllers

import (
	"coclone/models"
	"net/http"
)

const GridSize = 20

func HandleStartBattle(w http.ResponseWriter, r *http.Request) {
	state := &models.BattleState{
		Troops:     make(map[string]*models.Troop),
		Structures: make(map[string]*models.Structure),
		Grid:       make([][]*models.Cell, GridSize),
	}
	for i := 0; i < GridSize; i++ {
		state.Grid[i] = make([]*models.Cell, GridSize)

		for j := 0; j < GridSize; j++ {
			state.Grid[i][j] = &models.Cell{
				X:    i,
				Y:    j,
				Wall: false,
			}
		}
	}
	StartBattle(state)
}
