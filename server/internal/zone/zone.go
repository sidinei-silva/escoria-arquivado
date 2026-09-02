package zone

type ZoneID string

type ActionType string

const (
	ActionGather ActionType = "gather"
	ActionCombat ActionType = "combat"
	ActionCraft  ActionType = "craft"
	ActionRefine ActionType = "refine"
	ActionEquip  ActionType = "equip"
	ActionTravel ActionType = "travel"
)

var ValidActions = map[ActionType]bool{
	ActionGather: true,
	ActionCombat: true,
	ActionCraft:  true,
	ActionRefine: true,
	ActionEquip:  true,
	ActionTravel: true,
}

type Zone struct {
	ID            ZoneID
	Name          string
	Tier          int
	Actions       map[ActionType]bool
	Resources     []string
	AdjacentZones []ZoneID
}

func NewZone(id ZoneID, name string, tier int, actions []ActionType, resources []string, adjacentZones []ZoneID) *Zone {
	actionMap := make(map[ActionType]bool)
	for _, action := range actions {
		actionMap[action] = true
	}

	return &Zone{
		ID:            id,
		Name:          name,
		Tier:          tier,
		Actions:       actionMap,
		Resources:     resources,
		AdjacentZones: adjacentZones,
	}
}
