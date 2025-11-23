package services

import (
	"claimers/core/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

const DEFAULT_BOARD_SIZE int64 = 10

type GameService struct {
	Store models.Repository
}

func NewGameService(store models.Repository) *GameService {
	return &GameService{Store: store}
}

// NewGame creates a new game session and board, with
// an optional board size (default is DEFAULT_BOARD_SIZE)
// with the corresponding initialized tiles
func (s *GameService) NewGame(boardSize ...int) (models.GameSession, models.Board, []models.Tile, error) {
	b := DEFAULT_BOARD_SIZE
	if len(boardSize) > 0 {
		b = int64(boardSize[0])
	}
	board := &models.Board{
		ID:    uuid.New(),
		SizeX: null.IntFrom(b),
		SizeY: null.IntFrom(b),
	}
	err := s.Store.BoardCreate(board)
	if err != nil {
		return models.GameSession{}, models.Board{}, []models.Tile{}, fmt.Errorf("failed to create board: %w", err)
	}
	// create the tiles for the board
	tiles := make([]models.Tile, b*b)
	for x := int64(0); x < b; x++ {
		for y := int64(0); y < b; y++ {
			tiles = append(tiles, models.Tile{
				ID:       uuid.New(),
				BoardID:  board.ID,
				Position: models.Position{X: null.IntFrom(x), Y: null.IntFrom(y)},
			})
		}
	}
	err = s.Store.BoardTileCreateMany(tiles)
	if err != nil {
		return models.GameSession{}, models.Board{}, []models.Tile{}, fmt.Errorf("failed to create tiles: %w", err)
	}
	gameSession := &models.GameSession{
		ID:        uuid.New(),
		BoardID:   board.ID,
		StartTime: null.TimeFrom(time.Now()),
	}

	return *gameSession, *board, tiles, nil
}

func (s *GameService) CreatePlayer(gamerName string) (models.Player, error) {
	player := &models.Player{
		ID:        uuid.New(),
		GamerName: null.StringFrom(gamerName),
	}
	err := s.Store.PlayerCreate(player)
	if err != nil {
		return models.Player{}, fmt.Errorf("failed to create player: %w", err)
	}

	return *player, nil
}

func (s *GameService) AddPlayerToGameSession(gameSessionID uuid.UUID, playerID uuid.UUID) error {
	return s.Store.GameSessionAddPlayer(gameSessionID, playerID)
}

func (s *GameService) GetGameSession(id uuid.UUID) (*models.GameSession, error) {
	return s.Store.GameSessionGet(id)
}
