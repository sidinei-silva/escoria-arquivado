package gamedata

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

func loadZones(path string) (*ZonesFile, error) {
	return loadJSON[ZonesFile](path)
}
