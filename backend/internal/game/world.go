package game

import "fmt"

type World struct {
	Zones map[ZoneID]*Zone
}

func NewWorld() *World {
	return &World{
		Zones: make(map[ZoneID]*Zone),
	}
}

func (w *World) AddZone(zone *Zone) error {
	if zone == nil {
		return fmt.Errorf("zone is nil")
	}

	if _, exists := w.Zones[zone.ID]; exists {
		return fmt.Errorf("zona %q duplicada", zone.ID)
	}

	w.Zones[zone.ID] = zone

	return nil
}

func (w *World) Validate() error {
	for _, zone := range w.Zones {
		for _, adjacent := range zone.AdjacentZones {
			if _, exists := w.Zones[adjacent]; !exists {
				return fmt.Errorf(
					"zona %s: adjacente %q não existe",
					zone.ID,
					adjacent,
				)
			}
		}
	}

	return nil
}
