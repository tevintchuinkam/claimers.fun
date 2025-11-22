package temp

import (
	"claimers/models"
	"fmt"

	"github.com/google/uuid"
)

type TempStorage struct {
	games       map[uuid.UUID]*models.GameSession
	gamePlayers map[uuid.UUID][]*models.Player
	players     map[uuid.UUID]*models.Player
	boards      map[uuid.UUID]*models.Board
	boardTiles  map[uuid.UUID][]*models.Tile
}

// ensure TempStorage implements all the repository interfaces
var _ models.GameSessionRepository = (*TempStorage)(nil)

//var _ models.PlayerRepository = (*TempStorage)(nil)
//var _ models.BoardRepository = (*TempStorage)(nil)
//var _ models.BoardTilesRepository = (*TempStorage)(nil)

// GameSessionRepository methods
func (t *TempStorage) CreateGameSession(gameSession *models.GameSession) error {
	t.games[gameSession.ID] = gameSession
	return nil
}
func (t *TempStorage) GetGameSession(id uuid.UUID) (*models.GameSession, error) {
	return t.games[id], nil
}
func (t *TempStorage) UpdateGameSession(id uuid.UUID, update func(gameSession *models.GameSession) (models.GameSession, error)) (models.GameSession, error) {
	gameSession, ok := t.games[id]
	if !ok {
		return models.GameSession{}, fmt.Errorf("game session not found")
	}
	return update(gameSession)
}
func (t *TempStorage) DeleteGameSession(id uuid.UUID) error {
	delete(t.games, id)
	return nil
}
func (t *TempStorage) GetGameSessionPlayers(id uuid.UUID) ([]models.Player, error) {
	gamePlayers, ok := t.gamePlayers[id]
	if !ok {
		return nil, fmt.Errorf("game session players not found")
	}
	players := make([]models.Player, len(gamePlayers))
	for i, player := range gamePlayers {
		players[i] = *player
	}
	return players, nil
}
func (t *TempStorage) GetGameSessionBoard(id uuid.UUID) (*models.Board, error) {
	gameSession, ok := t.games[id]
	if !ok {
		return nil, fmt.Errorf("game session not found")
	}
	return t.boards[gameSession.BoardID], nil
}

// PlayerRepository methods
func (t *TempStorage) CreatePlayer(player *models.Player) error {
	t.players[player.ID] = player
	return nil
}
