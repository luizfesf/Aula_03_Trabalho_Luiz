package repository

import (
	"database/sql"
	"errors"

	"RPG_AULA03/internal/entity"
)

type EnemyRepository struct {
	db *sql.DB
}

func NewEnemyRepository(db *sql.DB) *EnemyRepository {
	return &EnemyRepository{db: db}
}

func (er *EnemyRepository) LoadEnemies() ([]*entity.Enemy, error) {
	rows, err := er.db.Query("SELECT id, nickname, life, attack FROM enemy")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enemies []*entity.Enemy
	for rows.Next() {
		var enemy entity.Enemy
		if err := rows.Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.Attack); err != nil {
			return nil, err
		}
		enemies = append(enemies, &enemy)
	}
	return enemies, nil
}

func (er *EnemyRepository) LoadEnemyById(id string) (*entity.Enemy, error) {
	var enemy entity.Enemy
	err := er.db.QueryRow("SELECT id, nickname, life, attack FROM enemy WHERE id = $1", id).Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.Attack)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &enemy, nil
}

func (er *EnemyRepository) LoadEnemyByNickname(nickname string) (*entity.Enemy, error) {
	var enemy entity.Enemy
	err := er.db.QueryRow("SELECT id, nickname, life, attack FROM enemy WHERE nickname LIKE $1", nickname).Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.Attack)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &enemy, nil
}

func (er *EnemyRepository) AddEnemy(enemy *entity.Enemy) (string, error) {
	_, err := er.db.Exec("INSERT INTO enemy (id, nickname, life, attack) VALUES ($1, $2, $3, $4)", enemy.ID, enemy.Nickname, enemy.Life, enemy.Attack)
	if err != nil {
		return "", err
	}
	return enemy.ID, nil
}

func (er *EnemyRepository) DeleteEnemyById(id string) error {
	_, err := er.db.Exec("DELETE FROM enemy WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (er *EnemyRepository) SaveEnemy(id string, enemy *entity.Enemy) error {
	_, err := er.db.Exec("UPDATE enemy SET nickname = $1, life = $2, attack = $3 WHERE id = $4", enemy.Nickname, enemy.Life, enemy.Attack, id)
	if err != nil {
		return err
	}
	return nil
}
