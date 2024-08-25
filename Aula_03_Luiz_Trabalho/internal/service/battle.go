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
    originalPlayerAttack := player.Attack
    originalEnemyAttack := enemy.Attack
    criticalMessage := ""

    if dice == 6 {
        player.Attack *= 2
        criticalMessage = "Critical Hit! Ataque do Jogador dobrado!"
    } else if dice == 1 {
        enemy.Attack *= 2
        criticalMessage = "Critical Hit! Ataque do Inimigo dobrado!"
    }

    var result string

    if dice <= 3 {
        // Inimigo ataca o jogador
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

        // Exibir resultados do ataque e status de vida/defesa/ataque
        Danos := "Inimigo atacou. Dano causado: " + strconv.Itoa(damage) +
            " | Vida do Jogador: " + strconv.Itoa(player.Life) +
            " | Defesa do Jogador: " + strconv.Itoa(player.Defesa) +
            " | Ataque do Inimigo: " + strconv.Itoa(enemy.Attack)

        // Result para os dados do inimigo
        enemyResult := "Dados do Inimigo: Vida: " + strconv.Itoa(enemy.Life) +
            " | Defesa: " + strconv.Itoa(enemy.Defesa) +
            " | Ataque: " + strconv.Itoa(enemy.Attack)

        result = Danos + "\n" + enemyResult + "\n" + criticalMessage

    } else {
        // Jogador ataca o inimigo
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

        // Exibir resultados do ataque e status de vida/defesa/ataque
        Danos := "Jogador atacou. Dano causado: " + strconv.Itoa(damage) +
            " | Vida do Inimigo: " + strconv.Itoa(enemy.Life) +
            " | Defesa do Inimigo: " + strconv.Itoa(enemy.Defesa) +
            " | Ataque do Jogador: " + strconv.Itoa(player.Attack)

        // Result para os dados do jogador
        playerResult := "Dados do Jogador: Vida: " + strconv.Itoa(player.Life) +
            " | Defesa: " + strconv.Itoa(player.Defesa) +
            " | Ataque: " + strconv.Itoa(player.Attack)

        result = Danos + "\n" + playerResult + "\n" + criticalMessage
    }

    // Verificar se alguém venceu a batalha
    if player.Life == 0 {
        battle.Result = "Inimigo venceu"
        result = "Inimigo venceu a batalha\n" + result
    } else if enemy.Life == 0 {
        battle.Result = "Jogador venceu"
        result = "Jogador venceu a batalha\n" + result
    } else {
        battle.Result = "A batalha continua"
    }

    if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
        return nil, "", err
    }

    // Restaurar os valores originais de ataque após o cálculo do dano
    player.Attack = originalPlayerAttack
    enemy.Attack = originalEnemyAttack

    return battle, result, nil
}
