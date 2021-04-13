package strategy

type WalkAction struct {
	walkVariants []int
}

func (action *WalkAction) GetActionDescription() string {
	return "Allows soldier to walk somewhere"
}

func (action *WalkAction) GetActionCost(actionVariant int) int {
	if action.walkVariants != nil && len(action.walkVariants) > actionVariant {
		return action.walkVariants[actionVariant]
	}
	return 0
}

func (action *WalkAction) GetActionVariantsNumber() int {
	return len(action.walkVariants)
}

func (action *WalkAction) IsActionPossible() bool {
	return false
}

func (action *WalkAction) PerformAction() {

}
