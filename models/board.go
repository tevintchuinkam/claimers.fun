package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Board struct {
	ID    uuid.UUID `db:"id"`
	SizeX null.Int  `db:"size_x"`
	SizeY null.Int  `db:"size_y"`
}

type BoardRepository interface {
	CreateBoard(board *Board) error
	GetBoard(id uuid.UUID) (*Board, error)
	GetBoardTiles(id uuid.UUID) ([]Tile, error)
	GetBoardTileAtPosition(id uuid.UUID, position Position) (*Tile, error)
	UpdateBoard(id uuid.UUID, update func(board *Board) (Board, error)) error
	DeleteBoard(id uuid.UUID) error
}
