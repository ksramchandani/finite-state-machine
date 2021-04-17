package main

type StateTransitionMap map[StateType]Transition

type Transition struct {
	Action              Action
	EventToNextStateMap EventToNextStateMap
}
