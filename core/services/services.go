package services

import (
	"claimers/core/models"
)

type GameService struct {
	Store models.Repository
}

func NewGameService(store models.Repository) *GameService {
	return &GameService{Store: store}
}
