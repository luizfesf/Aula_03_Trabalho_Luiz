package entity

import (
	"math/rand"
	"time"

	 "github.com/google/uuid"
)

type Battle struct {
	ID         string `json:"id"`
	PlayerID   string `json:"player_id"`
	EnemyID    string `json:"enemy_id"`
	PlayerName string `json:"player_name"`
	EnemyName  string `json:"enemy_name"`
	DiceThrown int    `json:"dice_thrown"`
	Result     string `json:"result"`
}

func NewBattle(playerID, enemyID, playerName, enemyName string) *Battle {
	rand.Seed(time.Now().UnixNano())
	return &Battle{
		ID:         uuid.New().String(),
		PlayerID:   playerID,
		EnemyID:    enemyID,
		PlayerName: playerName,
		EnemyName:  enemyName,
		DiceThrown: rand.Intn(6) + 1,
	}
}