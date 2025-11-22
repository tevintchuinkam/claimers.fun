package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Player struct {
	ID        uuid.UUID
	GamerName null.String
}

func (p *Player) Expand(boardID uuid.UUID, position Position) (Tile, error) {
	return Tile{}, nil
}

func (p *Player) Attack(boardID uuid.UUID, position Position) (Tile, error) {
	return Tile{}, nil
}

func (p *Player) Fortify(boardID uuid.UUID, position Position) (Tile, error) {
	return Tile{}, nil
}

type PlayerRepository interface {
	// CRUD actions
	CreatePlayer(player *Player) error
	GetPlayer(id uuid.UUID) (*Player, error)
	UpdatePlayer(id uuid.UUID, update func(player *Player) (Player, error)) error
	DeletePlayer(id uuid.UUID) error

	// Other actions
	GetPlayerActiveGameSession(playerID uuid.UUID) (*GameSession, error)
}
