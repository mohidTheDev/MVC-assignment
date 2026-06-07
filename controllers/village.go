package controllers

import (
	"coclone/models"
	"net/http"
)

func HandleStartBattle(w http.ResponseWriter, r *http.Request) {
	state := &models.BattleState{
		Troops:     make(map[string]*models.Troop),
		Structures: make(map[string]*models.Structure),
		Grid:       [][]*models.Cell{},
	}
	StartBattle(state)
}
