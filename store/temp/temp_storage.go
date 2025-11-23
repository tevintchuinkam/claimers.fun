package temp

import (
	"claimers/core/models"
	"fmt"

	"github.com/google/uuid"
)

type TempStorage struct {
	games       map[uuid.UUID]*models.GameSession
	gamePlayers map[uuid.UUID][]*models.Player
	players     map[uuid.UUID]*models.Player
	boards      map[uuid.UUID]*models.Board
	boardTiles  map[uuid.UUID][]*models.Tile
	tiles       map[uuid.UUID]*models.Tile // Store tiles by ID for BoardTileGet operations
}

// NewTempStorage creates a new TempStorage instance with initialized maps
func NewTempStorage() *TempStorage {
	return &TempStorage{
		games:       make(map[uuid.UUID]*models.GameSession),
		gamePlayers: make(map[uuid.UUID][]*models.Player),
		players:     make(map[uuid.UUID]*models.Player),
		boards:      make(map[uuid.UUID]*models.Board),
		boardTiles:  make(map[uuid.UUID][]*models.Tile),
		tiles:       make(map[uuid.UUID]*models.Tile),
	}
}

// ensure TempStorage implements all the repository interfaces
var _ models.Repository = (*TempStorage)(nil)

// GameSessionRepository methods
func (t *TempStorage) GameSessionCreate(gameSession *models.GameSession) (models.GameSession, error) {
	t.games[gameSession.ID] = gameSession
	return *gameSession, nil
}

func (t *TempStorage) GameSessionGet(id uuid.UUID) (*models.GameSession, error) {
	return t.games[id], nil
}

func (t *TempStorage) GameSessionUpdate(id uuid.UUID, update func(gameSession *models.GameSession) (models.GameSession, error)) (models.GameSession, error) {
	gameSession, ok := t.games[id]
	if !ok {
		return models.GameSession{}, fmt.Errorf("game session not found")
	}
	updated, err := update(gameSession)
	if err != nil {
		return models.GameSession{}, err
	}
	t.games[id] = &updated
	return updated, nil
}
func (t *TempStorage) GameSessionDelete(id uuid.UUID) error {
	delete(t.games, id)
	return nil
}
func (t *TempStorage) GameSessionGetPlayers(id uuid.UUID) ([]models.Player, error) {
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
func (t *TempStorage) GameSessionAddPlayer(gameSessionID uuid.UUID, playerID uuid.UUID) error {
	t.gamePlayers[gameSessionID] = append(t.gamePlayers[gameSessionID], &models.Player{ID: playerID})
	if _, ok := t.players[playerID]; !ok {
		return fmt.Errorf("player not found")
	}
	t.players[playerID] = &models.Player{ID: playerID}
	return nil
}

func (t *TempStorage) GameSessionGetBoard(id uuid.UUID) (*models.Board, error) {
	gameSession, ok := t.games[id]
	if !ok {
		return nil, fmt.Errorf("game session not found")
	}
	return t.boards[gameSession.BoardID], nil
}

// PlayerRepository methods
func (t *TempStorage) PlayerCreate(player *models.Player) error {
	t.players[player.ID] = player
	return nil
}

func (t *TempStorage) PlayerGet(id uuid.UUID) (*models.Player, error) {
	player, ok := t.players[id]
	if !ok {
		return nil, fmt.Errorf("player not found")
	}
	return player, nil
}

func (t *TempStorage) PlayerUpdate(id uuid.UUID, update func(player *models.Player) (models.Player, error)) error {
	player, ok := t.players[id]
	if !ok {
		return fmt.Errorf("player not found")
	}
	updated, err := update(player)
	if err != nil {
		return err
	}
	t.players[id] = &updated
	return nil
}

func (t *TempStorage) PlayerDelete(id uuid.UUID) error {
	delete(t.players, id)
	return nil
}

func (t *TempStorage) PlayerGetActiveGameSession(playerID uuid.UUID) (*models.GameSession, error) {
	// Find a game session where the player is a participant and the game hasn't ended
	for gameID, gameSession := range t.games {
		players, ok := t.gamePlayers[gameID]
		if !ok {
			continue
		}
		for _, player := range players {
			if player.ID == playerID {
				// Check if game has ended (EndTime is set)
				if !gameSession.EndTime.Valid {
					return gameSession, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("no active game session found for player")
}

// BoardRepository methods
func (t *TempStorage) BoardCreate(board *models.Board) error {
	t.boards[board.ID] = board
	return nil
}

func (t *TempStorage) BoardTileCreateMany(tiles []models.Tile) error {
	for _, tile := range tiles {
		t.tiles[tile.ID] = &tile
	}
	for _, tile := range tiles {
		if t.boardTiles[tile.BoardID] == nil {
			t.boardTiles[tile.BoardID] = []*models.Tile{}
		}
		t.boardTiles[tile.BoardID] = append(t.boardTiles[tile.BoardID], &tile)
	}
	return nil
}
func (t *TempStorage) BoardGet(boardID uuid.UUID) (*models.Board, error) {
	board, ok := t.boards[boardID]
	if !ok {
		return nil, fmt.Errorf("board not found")
	}
	return board, nil
}

func (t *TempStorage) BoardGetTiles(boardID uuid.UUID) ([]models.Tile, error) {
	tiles, ok := t.boardTiles[boardID]
	if !ok {
		return []models.Tile{}, nil
	}
	result := make([]models.Tile, len(tiles))
	for i, tile := range tiles {
		result[i] = *tile
	}
	return result, nil
}

func (t *TempStorage) BoardTileGetAtPosition(boardID uuid.UUID, position models.Position) (*models.Tile, error) {
	tiles, ok := t.boardTiles[boardID]
	if !ok {
		return nil, fmt.Errorf("board not found")
	}
	for _, tile := range tiles {
		if tile.Position.X == position.X && tile.Position.Y == position.Y {
			return tile, nil
		}
	}
	return nil, fmt.Errorf("tile not found at position")
}

func (t *TempStorage) BoardUpdate(boardID uuid.UUID, update func(board *models.Board) (models.Board, error)) error {
	board, ok := t.boards[boardID]
	if !ok {
		return fmt.Errorf("board not found")
	}
	updated, err := update(board)
	if err != nil {
		return err
	}
	t.boards[boardID] = &updated
	return nil
}

func (t *TempStorage) BoardDelete(boardID uuid.UUID) error {
	delete(t.boards, boardID)
	delete(t.boardTiles, boardID)
	// Also remove tiles from the tiles map
	for tileID, tile := range t.tiles {
		if tile.BoardID == boardID {
			delete(t.tiles, tileID)
		}
	}
	return nil
}

// BoardTileRepository methods
func (t *TempStorage) BoardTileCreate(tile *models.Tile) error {
	t.tiles[tile.ID] = tile
	// Also add to boardTiles map
	if t.boardTiles[tile.BoardID] == nil {
		t.boardTiles[tile.BoardID] = []*models.Tile{}
	}
	t.boardTiles[tile.BoardID] = append(t.boardTiles[tile.BoardID], tile)
	return nil
}

func (t *TempStorage) BoardTileGet(boardTileID uuid.UUID) (*models.Tile, error) {
	tile, ok := t.tiles[boardTileID]
	if !ok {
		return nil, fmt.Errorf("tile not found")
	}
	return tile, nil
}

func (t *TempStorage) BoardTileUpdate(boardTileID uuid.UUID, update func(tile *models.Tile) (models.Tile, error)) error {
	tile, ok := t.tiles[boardTileID]
	if !ok {
		return fmt.Errorf("tile not found")
	}
	updated, err := update(tile)
	if err != nil {
		return err
	}
	t.tiles[boardTileID] = &updated
	// Update in boardTiles map as well
	if tiles, ok := t.boardTiles[updated.BoardID]; ok {
		for i, tilePtr := range tiles {
			if tilePtr.ID == boardTileID {
				tiles[i] = &updated
				break
			}
		}
	}
	return nil
}

func (t *TempStorage) BoardTileDelete(boardTileID uuid.UUID) error {
	tile, ok := t.tiles[boardTileID]
	if !ok {
		return fmt.Errorf("tile not found")
	}
	delete(t.tiles, boardTileID)
	// Remove from boardTiles map
	if tiles, ok := t.boardTiles[tile.BoardID]; ok {
		for i, tilePtr := range tiles {
			if tilePtr.ID == boardTileID {
				t.boardTiles[tile.BoardID] = append(tiles[:i], tiles[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (t *TempStorage) BoardTileGetByOwnerID(ownerID uuid.UUID) ([]models.Tile, error) {
	var result []models.Tile
	for _, tile := range t.tiles {
		if tile.OwnerID == ownerID {
			result = append(result, *tile)
		}
	}
	return result, nil
}

func (t *TempStorage) BoardTileGetByBoardID(boardID uuid.UUID) ([]models.Tile, error) {
	tiles, ok := t.boardTiles[boardID]
	if !ok {
		return []models.Tile{}, nil
	}
	result := make([]models.Tile, len(tiles))
	for i, tile := range tiles {
		result[i] = *tile
	}
	return result, nil
}

func (t *TempStorage) BoardTileGetByPosition(position models.Position) ([]models.Tile, error) {
	var result []models.Tile
	for _, tile := range t.tiles {
		if tile.Position.X == position.X && tile.Position.Y == position.Y {
			result = append(result, *tile)
		}
	}
	return result, nil
}

func (t *TempStorage) BoardTileGetByPositionAndBoardID(position models.Position, boardID uuid.UUID) (models.Tile, error) {
	tiles, ok := t.boardTiles[boardID]
	if !ok {
		return models.Tile{}, fmt.Errorf("board not found")
	}
	for _, tile := range tiles {
		if tile.Position.X == position.X && tile.Position.Y == position.Y {
			return *tile, nil
		}
	}
	return models.Tile{}, fmt.Errorf("tile not found at position")
}
