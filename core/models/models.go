package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type GameSession struct {
	ID             uuid.UUID `db:"id"`
	BoardID        uuid.UUID `db:"board_id"`
	WinnerPlayerID uuid.UUID `db:"winner_player_id"`
	StartTime      null.Time `db:"start_time"`
	EndTime        null.Time `db:"end_time"`
}

type Board struct {
	ID    uuid.UUID `db:"id"`
	SizeX null.Int  `db:"size_x"`
	SizeY null.Int  `db:"size_y"`
}

func (b Board) Size() (int, int) {
	return int(b.SizeX.ValueOrZero()), int(b.SizeY.ValueOrZero())
}

type Player struct {
	ID        uuid.UUID
	GamerName null.String
}

type Tile struct {
	ID       uuid.UUID
	OwnerID  uuid.UUID
	Units    null.Int
	Position Position
	BoardID  uuid.UUID
}

type Position struct {
	X null.Int
	Y null.Int
}

type Repository interface {
	// GameSession
	GameSessionCreate(gameSession *GameSession) (GameSession, error)
	GameSessionGet(id uuid.UUID) (*GameSession, error)
	GameSessionUpdate(id uuid.UUID, update func(gameSession *GameSession) (GameSession, error)) (GameSession, error)
	GameSessionDelete(id uuid.UUID) error
	GameSessionGetPlayers(id uuid.UUID) ([]Player, error)
	GameSessionGetBoard(id uuid.UUID) (*Board, error)
	GameSessionAddPlayer(gameSessionID uuid.UUID, playerID uuid.UUID) error

	// Board
	BoardCreate(board *Board) error
	BoardGet(boardID uuid.UUID) (*Board, error)
	BoardGetTiles(boardID uuid.UUID) ([]Tile, error)
	BoardTileGetAtPosition(boardID uuid.UUID, position Position) (*Tile, error)
	BoardUpdate(boardID uuid.UUID, update func(board *Board) (Board, error)) error
	BoardDelete(boardID uuid.UUID) error

	// Player
	PlayerCreate(player *Player) error
	PlayerGet(id uuid.UUID) (*Player, error)
	PlayerUpdate(id uuid.UUID, update func(player *Player) (Player, error)) error
	PlayerDelete(id uuid.UUID) error
	PlayerGetActiveGameSession(playerID uuid.UUID) (*GameSession, error)

	// BoardTile
	BoardTileCreate(tile *Tile) error
	BoardTileCreateMany(tiles []Tile) error
	BoardTileGet(boardTileID uuid.UUID) (*Tile, error)
	BoardTileUpdate(boardTileID uuid.UUID, update func(tile *Tile) (Tile, error)) error
	BoardTileDelete(boardTileID uuid.UUID) error
	BoardTileGetByOwnerID(ownerID uuid.UUID) ([]Tile, error)
	BoardTileGetByBoardID(boardID uuid.UUID) ([]Tile, error)
	BoardTileGetByPosition(position Position) ([]Tile, error)
	BoardTileGetByPositionAndBoardID(position Position, boardID uuid.UUID) (Tile, error)
}
