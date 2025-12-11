package temp

import (
	"claimers/core/models"

	"github.com/google/uuid"
)

type TempStorage struct {
	players map[uuid.UUID]*models.Player
}

// NewTempStorage creates a new TempStorage instance with initialized maps
func NewTempStorage() *TempStorage {
	return &TempStorage{
		players: make(map[uuid.UUID]*models.Player),
	}
}
