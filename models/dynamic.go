package models

import "time"

type Player struct {
	Username            string    `json:"username" db:"username"`
	PWDHash             string    `json:"-" db:"pwd_hash"`
	AccountCreationDate time.Time `json:"account_creation_date" db:"account_creation_date"`
	LastLoginDate       time.Time `json:"last_login_date" db:"last_login_date"`
}

type Village struct {
	PlayerUsername string `json:"player_username" db:"player_username"`
	TownHallLevel  int    `json:"town_hall_level" db:"town_hall_level"`
	Gold           int    `json:"gold" db:"gold"`
	Elixir         int    `json:"elixir" db:"elixir"`
}

type Stats struct {
	Username        string `json:"username" db:"username"`
	AttacksWon      int    `json:"attacks_won" db:"attacks_won"`
	AttacksDefended int    `json:"attacks_defended" db:"attacks_defended"`
	GoldLooted      int    `json:"gold_looted" db:"gold_looted"`
	ElixirLooted    int    `json:"elixir_looted" db:"elixir_looted"`
	Trophies        int    `json:"trophies" db:"trophies"`
}

type PlacedBuilding struct {
	PlacementID      int64      `json:"placement_id" db:"placement_id"`
	PlayerUsername   string     `json:"player_username" db:"player_username"`
	PositionX        int        `json:"position_x" db:"position_x"`
	PositionY        int        `json:"position_y" db:"position_y"`
	Name             string     `json:"name" db:"name"`
	Level            int        `json:"level" db:"level"`
	UpgradeStartTime *time.Time `json:"upgrade_start_time" db:"upgrade_start_time"`
}

type Army struct {
	ArmyID         int64  `json:"army_id" db:"army_id"`
	PlayerUsername string `json:"player_username" db:"player_username"`
	Name           string `json:"name" db:"name"`
	Level          int    `json:"level" db:"level"`
	Quantity       int    `json:"quantity" db:"quantity"`
}

type BattleLog struct {
	BattleID int64  `json:"battle_id" db:"battle_id"`
	Attacker string `json:"attacker" db:"attacker"`
	Defender string `json:"defender" db:"defender"`
	Result   string `json:"result" db:"result"`
}

type BattleEvent struct {
	BattleID  int64     `json:"battle_id" db:"battle_id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Subject   string    `json:"subject" db:"subject"`
	Object    string    `json:"object" db:"object"`
	Action    string    `json:"action" db:"action"`
}
