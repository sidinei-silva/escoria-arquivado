package game

type ZoneID string
type ResourceID string

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
	Resources     []ResourceID
	AdjacentZones []ZoneID
}
