package gamedata

import (
	"errors"
	"escoria/internal/game"
	"fmt"
	"path/filepath"
)

type GameData struct {
	Zones map[game.ZoneID]*game.Zone
}

func Load(dir string) (*GameData, error) {
	zones, err := loadZones(filepath.Join(dir, "zones.json"))
	if err != nil {
		return nil, err
	}

	gd := &GameData{Zones: zones}

	if err := gd.validate(); err != nil {
		return nil, err
	}
	return gd, nil
}

func (gd *GameData) validate() error {
	var problems []error

	for _, z := range gd.Zones {
		for _, adj := range z.AdjacentZones {
			if _, ok := gd.Zones[adj]; !ok {
				problems = append(problems, fmt.Errorf(
					"zona %s: adjacente %q não existe", z.ID, adj))
			}
		}
	}

	return errors.Join(problems...)
}
