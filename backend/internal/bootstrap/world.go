package bootstrap

import (
	"fmt"

	"escoria/internal/game"
	"escoria/internal/gamedata"
)

func BuildWorld(data *gamedata.ZonesFile) (*game.World, error) {
	world := game.NewWorld()

	for _, definition := range data.Zones {
		zone, err := buildZone(definition)
		if err != nil {
			return nil, err
		}

		if err := world.AddZone(zone); err != nil {
			return nil, err
		}
	}

	if err := world.Validate(); err != nil {
		return nil, err
	}

	return world, nil
}

func buildZone(definition gamedata.ZoneDefinition) (*game.Zone, error) {
	zone := &game.Zone{
		ID:   game.ZoneID(definition.ID),
		Name: definition.Name,
		Tier: definition.Tier,
	}

	for _, action := range definition.Actions {
		actionType := game.ActionType(action)

		if !game.ValidActions[actionType] {
			return nil, fmt.Errorf(
				"zona %s: ação desconhecida %q",
				definition.ID,
				action,
			)
		}

		zone.Actions = append(zone.Actions, actionType)
	}

	for _, resource := range definition.Resources {
		zone.Resources = append(
			zone.Resources,
			game.ResourceID(resource),
		)
	}

	for _, adjacent := range definition.AdjacentZones {
		zone.AdjacentZones = append(
			zone.AdjacentZones,
			game.ZoneID(adjacent),
		)
	}

	return zone, nil
}
