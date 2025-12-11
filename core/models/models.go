package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Player struct {
	ID        uuid.UUID
	Gold      null.Int
	Diamonds  null.Int
	GamerName null.String
	Email     null.String
}

// Maps are the game maps that exist. they are not persisted to the database.
// They describe the territory where a game can be played. Each map has
// Different territories that are connected to each other by roads of a certain
// difficulty (it takes more effort for troops to travel through these)
type Map struct {
	ID                  uuid.UUID
	Name                string
	Territories         []MapTerritory
	AdjacentTerritories map[uuid.UUID]map[uuid.UUID]MapRoad
}

type MapRoad struct {
	ID          uuid.UUID
	Name        string
	Description string
	Difficulty  int
}

// MapTerritories make up the maps. Territories have different features that
// make them easier or more difficult for certain troops to travel through and
// fight in. You might have hills, that balloons or dragons can easily traverse,
// but golems might have trouble. In the snow, the ice wizaard will feel at home,
// but the fire mage will be powerless.
type MapTerritory struct {
	ID          uuid.UUID
	MapID       uuid.UUID
	Name        string
	Description string
	Features    TerritoryFeatures
}

// TerritoryFeatures describes the characteristics of a territory that affect
// how different troop types interact with it. This allows for flexible
// combinations of terrain types, environmental conditions, and other factors.
type TerritoryFeatures struct {
	// Temperature affects temperature-sensitive troops (ice wizard, fire mage)
	// Range: -100 (frozen) to 100 (scorching)
	Temperature int

	// Elevation affects flying vs ground troops
	// 0 = sea level, positive = hills/mountains, negative = valleys
	Elevation int

	// TerrainType is a string identifier for the primary terrain
	// Examples: "plains", "hills", "mountains", "snow", "desert", "forest", "swamp"
	TerrainType string

	// AdditionalTags allows for multiple terrain characteristics
	// Examples: ["rocky", "forested", "frozen", "arid"]
	Tags []string
}

// TroopType represents a type of military unit in the game
type TroopType struct {
	ID          uuid.UUID
	Name        string
	Description string
	// MovementType: "ground", "flying", "aquatic", "amphibious"
	MovementType string
	// TemperaturePreference: preferred temperature range
	// If troop is outside this range, they get penalties
	TemperaturePreference TemperatureRange
	// TerrainModifiers: how this troop type is affected by different terrains
	// Key: terrain type or tag, Value: modifier (positive = bonus, negative = penalty)
	TerrainModifiers map[string]TerrainModifier
}

// TemperatureRange defines the preferred temperature range for a troop type
type TemperatureRange struct {
	Min int // Minimum preferred temperature
	Max int // Maximum preferred temperature
	// If temperature is outside this range, apply penalties
}

// TerrainModifier describes how a terrain feature affects a troop type
type TerrainModifier struct {
	// MovementModifier: affects traversal speed (0.0 = can't move, 1.0 = normal, 2.0 = double speed)
	MovementModifier float64
	// CombatModifier: affects combat effectiveness (0.0 = powerless, 1.0 = normal, 2.0 = double power)
	CombatModifier float64
}

// CalculateTerritoryModifiers determines how a territory affects a specific troop type
// Returns modifiers for movement and combat effectiveness
func (t *TroopType) CalculateTerritoryModifiers(territory *MapTerritory) TerrainModifier {
	// Start with neutral modifiers
	result := TerrainModifier{
		MovementModifier: 1.0,
		CombatModifier:   1.0,
	}

	// Apply terrain type modifier if exists
	if mod, ok := t.TerrainModifiers[territory.Features.TerrainType]; ok {
		result.MovementModifier *= mod.MovementModifier
		result.CombatModifier *= mod.CombatModifier
	}

	// Apply tag modifiers
	for _, tag := range territory.Features.Tags {
		if mod, ok := t.TerrainModifiers[tag]; ok {
			result.MovementModifier *= mod.MovementModifier
			result.CombatModifier *= mod.CombatModifier
		}
	}

	// Apply elevation modifier for flying vs ground troops
	switch t.MovementType {
	case "flying":
		// Flying troops get bonus in high elevation
		if territory.Features.Elevation > 50 {
			result.MovementModifier *= 1.0
		}
	case "ground":
		// Ground troops get penalty in high elevation
		if territory.Features.Elevation > 50 {
			result.MovementModifier *= 0.8
		}
	}

	// Apply temperature modifier
	if territory.Features.Temperature < t.TemperaturePreference.Min ||
		territory.Features.Temperature > t.TemperaturePreference.Max {
		// Outside preferred range - apply penalty
		// Calculate how far outside the range
		var tempDiff int
		if territory.Features.Temperature < t.TemperaturePreference.Min {
			tempDiff = t.TemperaturePreference.Min - territory.Features.Temperature
		} else {
			tempDiff = territory.Features.Temperature - t.TemperaturePreference.Max
		}

		// Apply penalty based on how far outside the range (max 50% penalty)
		penalty := float64(tempDiff) / 100.0
		if penalty > 0.5 {
			penalty = 0.5
		}
		result.MovementModifier *= (1.0 - penalty)
		result.CombatModifier *= (1.0 - penalty)
	}

	return result
}

type Repository interface {
	// Player
	PlayerCreate(player *Player) error
	PlayerGet(id uuid.UUID) (*Player, error)
	PlayerUpdate(id uuid.UUID, update func(player *Player) (Player, error)) error
	PlayerDelete(id uuid.UUID) error

	// Matchmaking
	MatchmakingQueueAdd(playerID uuid.UUID) error
	MatchmakingQueueRemove(playerID uuid.UUID) error
	MatchmakingQueueGetAll() ([]Player, error)
	MatchmakingQueueGetCount() (int, error)
	MatchmakingQueueRemoveMany(playerIDs []uuid.UUID) error
}
