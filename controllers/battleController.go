package controllers

import (
	"coclone/database"
	"coclone/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func StartBattle(state *models.BattleState) {
	//Fetch village layout and troops from DB
	initiateVillage(state)
	fmt.Println("Starting Battle")
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
				fmt.Println("Battle Over")
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
	structuresRemaining := false
	for _, structure := range state.Structures {
		if structure.Name != "Wall" {
			structuresRemaining = true
			break
		}
	}
	if structuresRemaining == false || len(state.Troops) == 0 {
		return true
	}
	return false
}

func initiateVillage(state *models.BattleState) {
	//Fix users for testing. Later change according to battle
	defendingUser := "mohid_test"

	//Modify query and models to include other attributes of defenses later
	structureQuery := `
		SELECT 
			pb.Placement_ID, pb.Name, pb.PositionX, pb.PositionY,
			bls.HP,
			COALESCE(d.Damage, 0) AS Damage,
			COALESCE(d.Range, 0) AS AttackRange
		FROM Placed_Buildings pb
		JOIN Buildings_Level_Specific bls 
			ON pb.Name = bls.Name AND pb.Level = bls.Level
		LEFT JOIN Defenses d 
			ON bls.Name = d.Name AND bls.Level = d.Level
		WHERE pb.Player_Username = $1
	`

	rows, err := database.DB.Query(structureQuery, defendingUser)
	if err != nil {
		log.Println("Error fetching village layout:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var x, y float64
		var hp, damage, attackRange int

		err := rows.Scan(&id, &name, &x, &y, &hp, &damage, &attackRange)
		if err != nil {
			log.Println("Error reading structure row:", err)
			continue
		}

		strID := strconv.Itoa(id)

		state.Structures[strID] = &models.Structure{
			ID:          strID,
			Name:        name,
			X:           x,
			Y:           y,
			HP:          hp,
			MaxHP:       hp,
			Damage:      damage,
			AttackRange: float64(attackRange),
		}
	}

	for _, structure := range state.Structures {
		if structure.Name == "Wall" {
			addWallToGrid(structure, state)
		}
	}

	//For testing, fixing the username here
	//For the sake of testing, we will add some troops here. In real implementation, troops will be added when player deploys them
	attackingUser := "mohid_test2"
	troopQuery := `
		SELECT a.Name, a.Quantity, cld.HP, cld.Damage, cld.Projectile_Range, c.Walkspeed, c.Attack_Type
		FROM Army a
		JOIN Characters c ON a.Name = c.Name
		JOIN Character_Level_Dependent cld ON a.Name = cld.Name AND a.Level = cld.Level
		WHERE a.Player_Username = $1
	`

	troopRows, err := database.DB.Query(troopQuery, attackingUser)
	if err != nil {
		log.Println("Error fetching army:", err)
		return
	}
	defer troopRows.Close()

	troopCounter := 1

	for troopRows.Next() {
		var name string
		var quantity, hp, damage, projRange, walkspeed int
		var attackType string

		err := troopRows.Scan(&name, &quantity, &hp, &damage, &projRange, &walkspeed, &attackType)
		if err != nil {
			log.Println("Error reading troop row:", err)
			continue
		}

		for i := 0; i < quantity; i++ {
			strID := fmt.Sprintf("t_%d", troopCounter)

			isMelee := true
			if attackType != "Melee" {
				isMelee = false
			}

			state.Troops[strID] = &models.Troop{
				ID:          strID,
				Name:        name,
				X:           float64(2*i + 1),
				Y:           float64(2*i + 1),
				HP:          hp,
				MaxHP:       hp,
				Damage:      damage,
				MoveSpeed:   float64(walkspeed) / 100.0,
				IsMelee:     isMelee,
				AttackRange: float64(projRange),
			}
			troopCounter++
		}
	}
}

func addWallToGrid(wall *models.Structure, state *models.BattleState) {
	state.Grid[int(wall.X)][int(wall.Y)].Wall = true
}

func removeWallFromGrid(wall *models.Structure, state *models.BattleState) {
	state.Grid[int(wall.X)][int(wall.Y)].Wall = false
}
