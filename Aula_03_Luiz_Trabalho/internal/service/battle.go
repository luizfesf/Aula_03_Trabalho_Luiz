package service

import (
	"errors"


	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/repository"
)

type BattleService struct {
	PlayerRepository repository.PlayerRepository
	EnemyRepository  repository.EnemyRepository
	BattleRepository repository.BattleRepository
}

func NewBattleService(playerRepo repository.PlayerRepository, enemyRepo repository.EnemyRepository, battleRepo repository.BattleRepository) *BattleService {
	return &BattleService{
		PlayerRepository: playerRepo,
		EnemyRepository:  enemyRepo,
		BattleRepository: battleRepo,
	}
}

func (bs *BattleService) CreateBattle(playerNickname, enemyNickname string) (*entity.Battle, error) {
	player, err := bs.PlayerRepository.LoadPlayerByNickname(playerNickname)
	if err != nil || player == nil {
		return nil, errors.New("player not found")
	}

	enemy, err := bs.EnemyRepository.LoadEnemyByNickname(enemyNickname)
	if err != nil || enemy == nil {
		return nil, errors.New("enemy not found")
	}

	if player.Life <= 0 || enemy.Life <= 0 {
		return nil, errors.New("both player and enemy must have life > 0 to battle")
	}

	battle := entity.NewBattle(player.ID, enemy.ID, player.Nickname, enemy.Nickname)
	dice := battle.DiceThrown

	if dice <= 3 {
		player.Life -= enemy.Attack
		if player.Life < 0 {
			player.Life = 0
		}
		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, errors.New("failed to update player life")
		}
		battle.Result = "Enemy won"
	} else {
		enemy.Life -= player.Attack
		if enemy.Life < 0 {
			enemy.Life = 0
		}
		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, errors.New("failed to update enemy life")
		}
		battle.Result = "Player won"
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		return nil, err
	}

	return battle, nil
}

func (bs *BattleService) LoadBattles() ([]*entity.Battle, error) {
	return bs.BattleRepository.LoadBattles()
}


















