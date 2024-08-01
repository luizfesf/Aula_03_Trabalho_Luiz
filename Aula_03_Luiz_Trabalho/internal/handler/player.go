package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/service"
)

type PlayerHandler struct {
	PlayerService *service.PlayerService
}

func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{PlayerService: playerService}
}

func (ph *PlayerHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var player entity.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := ph.PlayerService.AddPlayer(player.Nickname, player.Life, player.Attack)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (ph *PlayerHandler) LoadPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	players, err := ph.PlayerService.LoadPlayers()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if err := ph.PlayerService.DeletePlayer(id); err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}

func (ph *PlayerHandler) LoadPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	player, err := ph.PlayerService.LoadPlayer(id)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (ph *PlayerHandler) SavePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var player entity.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := ph.PlayerService.SavePlayer(id, player.Nickname, player.Life, player.Attack)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
