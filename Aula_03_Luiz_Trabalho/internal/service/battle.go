package service

import (
    "encoding/json"
	"errors"
	"fmt"
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

	// Guardar os valores originais de ataque
	originalPlayerAttack := player.Attack
	originalEnemyAttack := enemy.Attack

	criticalMessage := ""

	// Aplicar efeito crítico conforme o dado
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

		// Mensagem de ataque
		result = fmt.Sprintf(
			"Inimigo atacou!\n"+
				"Dano causado: %d\n", damage,
		)

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

		// Mensagem de ataque
		result = fmt.Sprintf(
			"Jogador atacou!\n"+
				"Dano causado: %d\n", damage,
		)
	}

	// Detalhes dos jogadores
	playerDetails := map[string]interface{}{
		"Vida":   player.Life,
		"Defesa": player.Defesa,
		"Ataque": player.Attack,
	}

	enemyDetails := map[string]interface{}{
		"Vida":   enemy.Life,
		"Defesa": enemy.Defesa,
		"Ataque": enemy.Attack,
	}

	// Adicionar mensagem crítica, se houver
	if criticalMessage != "" {
		result = fmt.Sprintf("%s\n%s", result, criticalMessage)
	}

	// Verificar se alguém venceu a batalha
	if player.Life == 0 {
		battle.Result = "Inimigo venceu"
		result = fmt.Sprintf("Inimigo venceu a batalha\n%s", result)
	} else if enemy.Life == 0 {
		battle.Result = "Jogador venceu"
		result = fmt.Sprintf("Jogador venceu a batalha\n%s", result)
	} else {
		battle.Result = "A batalha continua"
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		return nil, "", err
	}

	// Restaurar os valores originais de ataque após o cálculo do dano
	player.Attack = originalPlayerAttack
	enemy.Attack = originalEnemyAttack

	// Exibir os resultados e dados dos jogadores como JSON formatado
	printPrettyJSON(map[string]interface{}{
		"result":   result,
		"player":   playerDetails,
		"enemy":    enemyDetails,
		"battle":   battle,
	})

	return battle, result, nil
}

func printPrettyJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Erro ao formatar o JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
