package gamedata

import (
	"escoria/internal/game"
	"fmt"
)

type ZoneDefinition struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Tier          *int     `json:"tier"`
	Actions       []string `json:"actions"`
	Resources     []string `json:"resources"`
	AdjacentZones []string `json:"adjacent_zones"`
}

type ZonesFile struct {
	Zones []ZoneDefinition `json:"zones"`
}

func loadZones(path string) (map[game.ZoneID]*game.Zone, error) {
	file, err := loadJSON[ZonesFile](path)
	if err != nil {
		return nil, err
	}

	zones := make(map[game.ZoneID]*game.Zone, len(file.Zones))

	for _, def := range file.Zones {
		if def.Tier == nil {
			return nil, fmt.Errorf("zona %s: tier ausente", def.ID)
		}

		actions := make(map[game.ActionType]bool, len(def.Actions))
		for _, a := range def.Actions {
			at := game.ActionType(a)
			if !game.ValidActions[at] {
				return nil, fmt.Errorf("zona %s: ação desconhecida %q", def.ID, a)
			}
			actions[at] = true
		}

		resources := make([]game.ResourceID, len(def.Resources))
		for i, r := range def.Resources {
			resources[i] = game.ResourceID(r)
		}

		adjacent := make([]game.ZoneID, len(def.AdjacentZones))
		for i, adj := range def.AdjacentZones {
			adjacent[i] = game.ZoneID(adj)
		}

		id := game.ZoneID(def.ID)
		zones[id] = &game.Zone{
			ID:            id,
			Name:          def.Name,
			Tier:          *def.Tier,
			Actions:       actions,
			Resources:     resources,
			AdjacentZones: adjacent,
		}
	}

	return zones, nil
}
