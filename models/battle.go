package models

type Troop struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	HP          int     `json:"hp"`
	MaxHP       int     `json:"max_hp"`
	Damage      int     `json:"damage"`
	MoveSpeed   float64 `json:"moveSpeed"`
	IsMelee     bool    `json:"is_melee"`
	AttackRange float64 `json:"attack_range"`
	MovePath    []*Cell `json:"move_path"`
	TargetID    string  `json:"target_id"`
}

type Structure struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	HP          int     `json:"hp"`
	MaxHP       int     `json:"max_hp"`
	Damage      int     `json:"damage"`
	AttackRange float64 `json:"attack_range"`
	TargetID    string  `json:"target_id"`
}

type BattleState struct {
	Troops     map[string]*Troop     `json:"troops"`
	Structures map[string]*Structure `json:"buildings"`
	Grid       [][]*Cell             `json:"grid"`
}

type Cell struct {
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Wall        bool   `json:"wall"`
	StructureID string `json:"structure_id,omitempty"`
}
