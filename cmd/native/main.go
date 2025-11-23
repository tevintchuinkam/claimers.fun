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
	windowWidth  = 1024
	windowHeight = 768
)

type Game struct {
	Session models.GameSession
	Board   models.Board
	Tiles   []models.Tile
	Players []models.Player
	Image   *ebiten.Image
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

	screen.DrawImage(g.Image, op)
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

	game := Game{
		Session: gameSession,
		Board:   board,
		Tiles:   tiles,
		Players: []models.Player{player1, player2},
	}

	return game, nil
}
