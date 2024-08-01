package entity

import "github.com/google/uuid"

type Player struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
}

func NewPlayer(nickname string, life, attack int) *Player {
	return &Player{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		Attack:   attack,
	}
}
