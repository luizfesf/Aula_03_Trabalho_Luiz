package entity

import "github.com/google/uuid"

type Player struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
	Defesa   int
}

func NewPlayer(nickname string, life, attack, defesa int) *Player {
	return &Player{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		Attack:   attack,
		Defesa:   defesa,
	}
}