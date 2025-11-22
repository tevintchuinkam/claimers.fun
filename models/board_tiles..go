package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Tile struct {
	OwnerID  uuid.UUID
	Units    null.Int
	Position Position
	BoardID  uuid.UUID
}

type BoardTilesRepository interface {
	// CRUD actions
	CreateBoardTile(tile *Tile) error
	GetBoardTile(id uuid.UUID) (*Tile, error)
	UpdateBoardTile(id uuid.UUID, update func(tile *Tile) (Tile, error)) error
	DeleteBoardTile(id uuid.UUID) error

	// Other actions
	GetBoardTileAtPosition(boardID uuid.UUID, position Position) (*Tile, error)
	GetBoardTilesByOwnerID(ownerID uuid.UUID) ([]Tile, error)
	GetBoardTilesByBoardID(boardID uuid.UUID) ([]Tile, error)
	GetBoardTilesByPosition(position Position) ([]Tile, error)
	GetBoardTilesByPositionAndBoardID(position Position, boardID uuid.UUID) ([]Tile, error)
	GetBoardTilesByPositionAndOwnerID(position Position, ownerID uuid.UUID) ([]Tile, error)
}
