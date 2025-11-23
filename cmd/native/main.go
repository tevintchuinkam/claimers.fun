package main

import (
	"claimers/core/services"
	"claimers/store/temp"
	"log"
)

func main() {
	// create a new game session
	store := temp.NewTempStorage()

	gameService := services.NewGameService(store)

	game, err := gameService.NewGame()
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

	err = gameService.AddPlayerToGameSession(game.ID, player1.ID)
	if err != nil {
		log.Fatalf("failed to add player to game session: %v", err)
	}
	err = gameService.AddPlayerToGameSession(game.ID, player2.ID)
	if err != nil {
		log.Fatalf("failed to add player to game session: %v", err)
	}

}
