package models

import (
	"time"

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

func (g *GameSession) Start() {
	g.StartTime = null.TimeFrom(time.Now())
}

func (g *GameSession) End() {
	g.EndTime = null.TimeFrom(time.Now())
}

func (g *GameSession) HasWinner() (bool, error) {
	return g.WinnerPlayerID != uuid.Nil, nil
}

type GameSessionRepository interface {
	// CRUD actions
	CreateGameSession(gameSession *GameSession) error
	GetGameSession(id uuid.UUID) (*GameSession, error)
	UpdateGameSession(id uuid.UUID, update func(gameSession *GameSession) (GameSession, error)) (GameSession, error)
	DeleteGameSession(id uuid.UUID) error

	// Other actions
	GetGameSessionPlayers(id uuid.UUID) ([]Player, error)
	GetGameSessionBoard(id uuid.UUID) (*Board, error)
}
