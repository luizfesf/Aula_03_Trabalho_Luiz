package service

import (
	"errors"
	"strconv"

	"AULA_03_LUIZ_TRABALHO/internal/entity"
	"AULA_03_LUIZ_TRABALHO/internal/repository"
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

func (bs *BattleService) CreateBattle(playerNickname, enemyNickname string) (*entity.Battle, string, error) {
	player, err := bs.PlayerRepository.LoadPlayerByNickname(playerNickname)
	if err != nil || player == nil {
		return nil, "", errors.New("jogador não encontrado")
	}

	enemy, err := bs.EnemyRepository.LoadEnemyByNickname(enemyNickname)
	if err != nil || enemy == nil {
		return nil, "", errors.New("inimigo não encontrado")
	}

	if player.Life <= 0 || enemy.Life <= 0 {
		return nil, "", errors.New("tanto o jogador quanto o inimigo devem ter vida > 0 para batalhar")
	}

	battle := entity.NewBattle(player.ID, enemy.ID, player.Nickname, enemy.Nickname)
	dice := battle.DiceThrown

	// Implementação da funcionalidade "Critical"
	if dice == 6 {
		player.Attack *= 2
	} else if dice == 1 {
		enemy.Attack *= 2
	}

	var result string

	if dice <= 3 {
		damage := enemy.Attack - player.Defesa
		if damage < 0 {
			damage = 0
		}
		player.Life -= damage
		if player.Life < 0 {
			player.Life = 0
		}
		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do jogador")
		}
		Danos := "Inimigo atacou. Dano causado: " + strconv.Itoa(damage)

		result = Danos
	} else {
		// Calcular o dano considerando a defesa
		damage := player.Attack - enemy.Defesa
		if damage < 0 {
			damage = 0
		}
		enemy.Life -= damage
		if enemy.Life < 0 {
			enemy.Life = 0
		}
		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do inimigo")
		}

		// Result para o dano causado e dados do inimigo
		Danos := "Jogador atacou. Dano causado: " + strconv.Itoa(damage) +
			" | Vida do Inimigo: " + strconv.Itoa(enemy.Life) +
			" | Defesa do Inimigo: " + strconv.Itoa(enemy.Defesa) +
			" | Ataque do Jogador: " + strconv.Itoa(player.Attack)

		// Result para os dados do jogador
		playerResult := "Dados do Jogador: Vida: " + strconv.Itoa(player.Life) +
			" | Defesa: " + strconv.Itoa(player.Defesa) +
			" | Ataque: " + strconv.Itoa(player.Attack)

		result = Danos + "\n" + playerResult
	}

	if player.Life == 0 {
		battle.Result = "Inimigo venceu"
		result = "Inimigo venceu a batalha"
	} else if enemy.Life == 0 {
		battle.Result = "Jogador venceu"
		result = "Jogador venceu a batalha"
	} else {
		battle.Result = "A batalha continua"
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		return nil, "", err
	}

	return battle, result, nil
}
