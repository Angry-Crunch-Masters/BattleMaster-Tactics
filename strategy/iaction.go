package strategy

//IAction defines action for strategy unit
type IAction interface {
	GetActionDescription() string
	GetActionCost(actionVariant int) int
	GetActionVariantsNumber() int
	IsActionPossible(entity IStrategyEntity) bool
	PerformAction(entity IStrategyEntity)
}
