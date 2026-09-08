package gamedata

import (
	"path/filepath"
)

type GameData struct {
	Zones *ZonesFile
}

func Load(dir string) (*GameData, error) {
	zones, err := LoadZones(filepath.Join(dir, "zones.json"))
	if err != nil {
		return nil, err
	}

	return &GameData{
		Zones: zones,
	}, nil
}
