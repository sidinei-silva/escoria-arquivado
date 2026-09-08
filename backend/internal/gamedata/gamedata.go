package gamedata

import (
	"log/slog"
	"path/filepath"
)

type GameData struct {
	Zones *ZonesFile
	// Dialogs    *DialogsFile
	// Objectives *ObjectivesFile
	// Items      *ItemsFile
	// Recipes    *RecipesFile
}

func Load(dir string) (*GameData, error) {
	zones, err := loadZones(filepath.Join(dir, "zones.json"))
	if err != nil {
		slog.Error("Failed to load zones", "err", err)
		return nil, err
	}

	return &GameData{
		Zones: zones,
	}, nil
}
