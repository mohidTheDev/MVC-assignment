package models

type Building struct {
	Name          string `json:"name" db:"name"`
	Type          string `json:"type" db:"type"`
	BuildResource string `json:"build_resource" db:"build_resource"` //gold/elixir
	SizeX         int    `json:"size_x" db:"size_x"`
	SizeY         int    `json:"size_y" db:"size_y"`
}

type BuildingLevelSpecific struct {
	Name             string `json:"name" db:"name"`
	Level            int    `json:"level" db:"level"`
	MinTownHallLevel int    `json:"min_town_hall_level" db:"min_town_hall_level"`
	BuildCost        int    `json:"build_cost" db:"build_cost"`
	HP               int    `json:"hp" db:"hp"`
	UpgradeTime      int    `json:"upgrade_time" db:"upgrade_time"`
}

type Character struct {
	Name           string `json:"name" db:"name"`
	HousingSpace   int    `json:"housing_space" db:"housing_space"`
	Walkspeed      int    `json:"walkspeed" db:"walkspeed"`
	MovementType   string `json:"movement_type" db:"movement_type"` //ground/air
	AttackType     string `json:"attack_type" db:"attack_type"`     //melee/ranged
	ProjectileType string `json:"projectile_type" db:"projectile_type"`
}

type CharacterLevelDependent struct {
	Name             string `json:"name" db:"name"`
	Level            int    `json:"level" db:"level"`
	UpgradeCost      int    `json:"upgrade_cost" db:"upgrade_cost"`
	MinBarracksLevel int    `json:"min_barracks_level" db:"min_barracks_level"`
	HP               int    `json:"hp" db:"hp"`
	Damage           int    `json:"damage" db:"damage"`
	ProjectileRange  int    `json:"projectile_range" db:"projectile_range"`
}
