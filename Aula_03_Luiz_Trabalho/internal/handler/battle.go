package handler

import (
	"encoding/json"
	"net/http"
    "RPG_AULA03/internal/service"
)

type BattleHandler struct {
	BattleService *service.BattleService
}

func NewBattleHandler(battleService *service.BattleService) *BattleHandler {
	return &BattleHandler{BattleService: battleService}
}

func (bh *BattleHandler) CreateBattle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		Player string `json:"player"`
		Enemy  string `json:"enemy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		return
	}

	battle, err := bh.BattleService.CreateBattle(request.Player, request.Enemy)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(battle)
}

func (bh *BattleHandler) LoadBattles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	battles, err := bh.BattleService.LoadBattles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(battles)
}
