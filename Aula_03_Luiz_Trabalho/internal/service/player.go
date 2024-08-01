package service

import (
	"errors"
	"fmt"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/repository"
)

type PlayerService struct {
	PlayerRepository repository.PlayerRepository
}

func NewPlayerService(PlayerRepository repository.PlayerRepository) *PlayerService {
	return &PlayerService{PlayerRepository: PlayerRepository}
}

func (ps *PlayerService) AddPlayer(nickname string, life, attack int) (*entity.Player, error) {
	if nickname == "" || life == 0 || attack == 0 {
		return nil, errors.New("player nickname, life and attack is required")
	}

	if len(nickname) > 255 {
		return nil, errors.New("player nickname cannot exceed 255 characters")
	}

	if attack > 10 || attack <= 0 {
		return nil, errors.New("player attack must be between 1 and 10")
	}

	if life > 100 || life <= 0 {
		return nil, errors.New("player life must be between 1 and 100")
	}

	player, err := ps.PlayerRepository.LoadPlayerByNickname(nickname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player != nil {
		return nil, errors.New("player nickname already exits")
	}

	player = entity.NewPlayer(nickname, life, attack)
	if _, err := ps.PlayerRepository.AddPlayer(player); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return player, nil
}

func (ps *PlayerService) LoadPlayers() ([]*entity.Player, error) {
	players, err := ps.PlayerRepository.LoadPlayers()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if players == nil {
		return []*entity.Player{}, nil
	}
	return players, nil
}

func (ps *PlayerService) DeletePlayer(id string) error {
	player, err := ps.PlayerRepository.LoadPlayerById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if player == nil {
		return errors.New("player id not found")
	}
	if err := ps.PlayerRepository.DeletePlayerById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (ps *PlayerService) LoadPlayer(id string) (*entity.Player, error) {
	player, err := ps.PlayerRepository.LoadPlayerById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player id not found")
	}
	return player, nil
}

func (ps *PlayerService) SavePlayer(id, nickname string, life, attack int) (*entity.Player, error) {
	player, err := ps.PlayerRepository.LoadPlayerById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player id not found")
	}

	if nickname != "" && nickname != player.Nickname {
		hasNickname, err := ps.PlayerRepository.LoadPlayerByNickname(nickname)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasNickname != nil {
			return nil, errors.New("player nickname already exits")
		}
		if len(nickname) > 255 {
			return nil, errors.New("player nickname cannot exceed 255 characters")
		}
		player.Nickname = nickname
	}

	if attack != 0 && attack != player.Attack {
		if attack > 10 || attack <= 0 {
			return nil, errors.New("player attack must be between 1 and 10")
		}
		player.Attack = attack
	}

	if life != 0 && life != player.Life {
		if life > 100 || life <= 0 {
			return nil, errors.New("player life must be between 1 and 100")
		}
		player.Life = life
	}

	if err := ps.PlayerRepository.SavePlayer(id, player); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return player, nil
}
