package dataloader

import (
	"encoding/json"
	"escoria/internal/zone"
	"fmt"
	"log/slog"
	"os"
)

type ZoneDefinition struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Tier          int      `json:"tier"`
	Actions       []string `json:"actions"`
	Resources     []string `json:"resources"`
	AdjacentZones []string `json:"adjacent_zones"`
}

type ZonesFile struct {
	Zones []ZoneDefinition `json:"zones"`
}

var zonePathFile = "../data/contents/zones.json"

func LoadZonesFromFile() ([]*zone.Zone, error) {
	jsonFile, err := os.ReadFile(zonePathFile)

	if err != nil {
		slog.Error("Falha ao carregar arquivo", "zonePathFile", zonePathFile, "err", err)
		return nil, err
	}

	var zonesFile ZonesFile
	err = json.Unmarshal(jsonFile, &zonesFile)
	if err != nil {
		slog.Error("Falha ao fazer unmarshal do arquivo", "zonePathFile", zonePathFile, "err", err)
		return nil, err
	}

	zones := make([]*zone.Zone, 0, len(zonesFile.Zones))

	for _, zoneDef := range zonesFile.Zones {
		actions := make([]zone.ActionType, len(zoneDef.Actions))

		for i, action := range zoneDef.Actions {
			if action == "" {
				slog.Error("Ação vazia encontrada", "zoneID", zoneDef.ID)
				return nil, fmt.Errorf("ação vazia encontrada na zona %s", zoneDef.ID)
			}

			if !zone.ValidActions[zone.ActionType(action)] {
				slog.Error("Ação inválida encontrada", "zoneID", zoneDef.ID, "action", action)
				os.Exit(1)
			}

			actions[i] = zone.ActionType(action)
		}

		adjacentZones := make([]zone.ZoneID, len(zoneDef.AdjacentZones))

		for i, adjZone := range zoneDef.AdjacentZones {

			// Verificar se a zona adjacente existe no arquivo de zonas
			zoneExists := false
			for _, z := range zonesFile.Zones {
				if z.ID == adjZone {
					zoneExists = true
					break
				}
			}

			if !zoneExists {
				slog.Error("Zona adjacente inválida encontrada", "zoneID", zoneDef.ID, "adjacentZone", adjZone)
				os.Exit(1)
			}

			adjacentZones[i] = zone.ZoneID(adjZone)
		}

		zone := zone.NewZone(zone.ZoneID(zoneDef.ID), zoneDef.Name, zoneDef.Tier, actions, zoneDef.Resources, adjacentZones)
		zones = append(zones, zone)
	}

	return zones, nil
}
