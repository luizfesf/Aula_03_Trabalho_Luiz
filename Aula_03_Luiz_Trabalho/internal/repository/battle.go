package repository

import (
	"AULA_03_LUIZ_TRABALHO/internal/entity"
	"database/sql"
)

type BattleRepository struct {
	db *sql.DB
}

func NewBattleRepository(db *sql.DB) *BattleRepository {
	return &BattleRepository{db: db}
}

func (br *BattleRepository) AddBattle(battle *entity.Battle) (string, error) {
	_, err := br.db.Exec(
		"INSERT INTO battle (id, playerid, enemyid, playername, enemyname, dicethrown, result) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		battle.ID, battle.PlayerID, battle.EnemyID, battle.PlayerName, battle.EnemyName, battle.DiceThrown, battle.Result,
	)
	if err != nil {
		return "", err
	}
	return battle.ID, nil
}

// Função para carregar batalhas do banco de dados
func (br *BattleRepository) LoadBattles() ([]*entity.Battle, error) {
	rows, err := br.db.Query("SELECT id, playerid, enemyid, playername, enemyname, dicethrown, result FROM battle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var battles []*entity.Battle
	for rows.Next() {
		var battle entity.Battle
		if err := rows.Scan(&battle.ID, &battle.PlayerID, &battle.EnemyID, &battle.PlayerName, &battle.EnemyName, &battle.DiceThrown, &battle.Result); err != nil {
			return nil, err
		}
		battles = append(battles, &battle)
	}
	return battles, nil
}