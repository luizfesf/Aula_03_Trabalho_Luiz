package service

import (
	"errors"
	"fmt"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/repository"
)

type EnemyService struct {
	EnemyRepository repository.EnemyRepository
}

func NewEnemyService(EnemyRepository repository.EnemyRepository) *EnemyService {
	return &EnemyService{EnemyRepository: EnemyRepository}
}

func (es *EnemyService) AddEnemy(nickname string, life, attack int) (*entity.Enemy, error) {
	if nickname == "" || life == 0 || attack == 0 {
		return nil, errors.New("enemy nickname, life and attack is required")
	}

	if len(nickname) > 255 {
		return nil, errors.New("enemy nickname cannot exceed 255 characters")
	}

	if attack > 10 || attack <= 0 {
		return nil, errors.New("enemy attack must be between 1 and 10")
	}

	if life > 100 || life <= 0 {
		return nil, errors.New("enemy life must be between 1 and 100")
	}

	enemy, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy != nil {
		return nil, errors.New("enemy nickname already exists")
	}

	enemy = entity.NewEnemy(nickname, life, attack)
	if _, err := es.EnemyRepository.AddEnemy(enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}

func (es *EnemyService) LoadEnemies() ([]*entity.Enemy, error) {
	enemies, err := es.EnemyRepository.LoadEnemies()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if enemies == nil {
		return []*entity.Enemy{}, nil
	}
	return enemies, nil
}

func (es *EnemyService) DeleteEnemy(id string) error {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if enemy == nil {
		return errors.New("enemy id not found")
	}
	if err := es.EnemyRepository.DeleteEnemyById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (es *EnemyService) LoadEnemy(id string) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}
	return enemy, nil
}

func (es *EnemyService) SaveEnemy(id, nickname string, life, attack int) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}

	if nickname != "" && nickname != enemy.Nickname {
		hasNickname, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasNickname != nil {
			return nil, errors.New("enemy nickname already exists")
		}
		if len(nickname) > 255 {
			return nil, errors.New("enemy nickname cannot exceed 255 characters")
		}
		enemy.Nickname = nickname
	}

	if attack != 0 && attack != enemy.Attack {
		if attack > 10 || attack <= 0 {
			return nil, errors.New("enemy attack must be between 1 and 10")
		}
		enemy.Attack = attack
	}

	if life != 0 && life != enemy.Life {
		if life > 100 || life <= 0 {
			return nil, errors.New("enemy life must be between 1 and 100")
		}
		enemy.Life = life
	}

	if err := es.EnemyRepository.SaveEnemy(id, enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}
