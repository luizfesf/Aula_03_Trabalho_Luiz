package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/service"
)

type EnemyHandler struct {
	EnemyService *service.EnemyService
}

func NewEnemyHandler(enemyService *service.EnemyService) *EnemyHandler {
	return &EnemyHandler{EnemyService: enemyService}
}

func (eh *EnemyHandler) AddEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var enemy entity.Enemy
	if err := json.NewDecoder(r.Body).Decode(&enemy); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := eh.EnemyService.AddEnemy(enemy.Nickname, enemy.Life, enemy.Attack)
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

func (eh *EnemyHandler) LoadEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enemies, err := eh.EnemyService.LoadEnemies()
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
	json.NewEncoder(w).Encode(enemies)
}

func (eh *EnemyHandler) DeleteEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if err := eh.EnemyService.DeleteEnemy(id); err != nil {
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

func (eh *EnemyHandler) LoadEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	enemy, err := eh.EnemyService.LoadEnemy(id)

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
	json.NewEncoder(w).Encode(enemy)
}

func (eh *EnemyHandler) SaveEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var enemy entity.Enemy
	if err := json.NewDecoder(r.Body).Decode(&enemy); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := eh.EnemyService.SaveEnemy(id, enemy.Nickname, enemy.Life, enemy.Attack)
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
