package main

import (
	"fmt"
	"sync"
)

type StateMachine struct {
	Previous           StateType
	Current            StateType
	StateTransitionMap StateTransitionMap
	mutex              sync.Mutex
}

func (s *StateMachine) GetNextState(event EventType) (StateType, error) {
	transition, ok := s.StateTransitionMap[s.Current]
	if !ok {
		return StateError, fmt.Errorf("No transition exists")
	}

	nextState, ok := transition.EventToNextStateMap[event]
	if !ok {
		return StateError, fmt.Errorf("No next state exists")
	}

	return nextState, nil
}

func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext, stateTransitions *[]StateType) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		*stateTransitions = append(*stateTransitions, s.Current)
		nextState, err := s.GetNextState(event)
		if err != nil {
			return err
		}

		transition, ok := s.StateTransitionMap[nextState]
		if !ok {
			return fmt.Errorf("No transition exists")
		}

		s.Previous = s.Current
		s.Current = nextState
		fmt.Printf("%v --> %v\n", s.Previous, s.Current)
		nextEvent := transition.Action.Execute(eventCtx)
		if nextEvent == EventNoOP {
			*stateTransitions = append(*stateTransitions, s.Current)
			return nil
		}

		event = nextEvent

	}
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		Current: StateInitial,
		StateTransitionMap: StateTransitionMap{

			StateInitial: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventOrderCreated: StateOrderCreated,
				},
			},

			StateOrderCreated: Transition{
				Action: new(ActionOrderCreated),
				EventToNextStateMap: EventToNextStateMap{
					EventOrderPlaced: StateOrderPlaced,
					EventOrderFailed: StateOrderFailed,
				},
			},

			StateOrderFailed: Transition{
				Action: new(ActionOrderFailed),
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},

			StateOrderPlaced: Transition{
				Action: new(ActionOrderPlaced),
				EventToNextStateMap: EventToNextStateMap{
					EventChargeCard: StateChargingCard,
				},
			},

			StateChargingCard: Transition{
				Action: new(ActionCharingCard),
				EventToNextStateMap: EventToNextStateMap{
					EventTransactionFailed: StateTransactionFailed,
					EventOrderShipped:      StateOrderShipped,
				},
			},

			StateTransactionFailed: Transition{
				Action: new(ActionTransactionFailed),
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},

			StateOrderShipped: Transition{
				Action: new(ActionOrderShipped),
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},
		},
	}
}
