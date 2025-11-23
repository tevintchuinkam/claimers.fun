package main

import (
	"claimers/core/design"
	"claimers/core/models"
	"claimers/core/services"
	"claimers/store/temp"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth     int     = 1024
	windowHeight    int     = 768
	tileMarginRatio float64 = 0.05
)

type Game struct {
	Session models.GameSession
	Board   models.Board
	Tiles   []Tile
	Players []models.Player
	Image   *ebiten.Image
}

type Tile struct {
	BoardSizeX int64
	BoardSizeY int64
	Tile       models.Tile
}

func (t Tile) Update() error {
	return nil
}

func (t Tile) Position() (int, int) {
	return int(t.Tile.Position.X.ValueOrZero()), int(t.Tile.Position.Y.ValueOrZero())
}

func (t Tile) Draw(screen *ebiten.Image) {
	i, j := t.Position()
	tileWidth := (float64(windowWidth) / float64(t.BoardSizeX))
	tileHeight := (float64(windowHeight) / float64(t.BoardSizeY))
	op := &ebiten.DrawImageOptions{}
	x := int(float64(i) * tileWidth)
	y := int(float64(j) * tileHeight)
	op.GeoM.Translate(float64(x), float64(y))
	tileImage := ebiten.NewImage(int(tileWidth), int(tileHeight))
	tileImage.Fill(design.TileColor)
	screen.DrawImage(tileImage, op)
}

func main() {
	// create a new game session
	store := temp.NewTempStorage()

	gameService := services.NewGameService(store)

	game, err := NewGame(gameService)
	if err != nil {
		log.Fatalf("failed to create game: %v", err)
	}

	ebiten.SetWindowTitle("Claimers")
	ebiten.SetWindowSize(windowWidth, windowHeight)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to run game: %v", err)
	}
}

func (g Game) Layout(outisdeWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	if g.Image == nil {
		g.Image = ebiten.NewImage(g.Board.Size())
	}

	screen.Fill(design.BackgroundColor)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.Image.Bounds().Dx(), g.Image.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))

	// now we need to draw all the tiles onto the board
	for _, tile := range g.Tiles {
		tile.Draw(screen)
	}

}

func NewGame(gameService *services.GameService) (Game, error) {
	gameSession, board, tiles, err := gameService.NewGame()
	if err != nil {
		log.Fatalf("failed to create game: %v", err)
	}

	// create two players
	player1, err := gameService.CreatePlayer("Player 1")
	if err != nil {
		log.Fatalf("failed to create player: %v", err)
	}
	player2, err := gameService.CreatePlayer("Player 2")
	if err != nil {
		log.Fatalf("failed to create player: %v", err)
	}

	err = gameService.AddPlayerToGameSession(gameSession.ID, player1.ID)
	if err != nil {
		log.Fatalf("failed to add player to game session: %v", err)
	}
	err = gameService.AddPlayerToGameSession(gameSession.ID, player2.ID)
	if err != nil {
		log.Fatalf("failed to add player to game session: %v", err)
	}

	// now we need to create the game image
	tileObjects := make([]Tile, len(tiles))
	for i, tile := range tiles {
		tileObjects[i] = Tile{Tile: tile, BoardSizeX: board.SizeX.ValueOrZero(), BoardSizeY: board.SizeY.ValueOrZero()}
	}

	game := Game{
		Session: gameSession,
		Board:   board,
		Tiles:   tileObjects,
		Players: []models.Player{player1, player2},
	}

	return game, nil
}
