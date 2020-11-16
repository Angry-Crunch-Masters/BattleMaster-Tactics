package basic

//EntityType is enum for entity types
type EntityType int

const (
	//StrategyEntity is used as strategy type for entities manager
	StrategyEntity EntityType = iota
	//Other is used as other entity type for entites manager
	Other
)
