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

func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		fmt.Printf("In Current state = %v\n", s.Current)
		nextState, err := s.GetNextState(event)
		if err != nil {
			return err
		}
		fmt.Printf("Got next state = %v\n", nextState)

		transition, ok := s.StateTransitionMap[nextState]
		if !ok {
			return fmt.Errorf("No transition exists")
		}

		s.Previous = s.Current
		s.Current = nextState
		nextEvent := transition.Action.Execute(eventCtx)
		if nextEvent == EventNoOP {
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
				Action: new(ActionCreateOrder),
				EventToNextStateMap: EventToNextStateMap{
					EventOrderPlaced: StateOrderPlaced,
					EventOrderFailed: StateOrderFailed,
				},
			},

			StateOrderFailed: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},

			StateOrderPlaced: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventChargeCard: StateCharingCard,
				},
			},

			StateCharingCard: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventTransactionFailed: StateTransactionFailed,
					EventOrderShipped:      StateOrderShipped,
				},
			},

			StateTransactionFailed: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},

			StateOrderShipped: Transition{
				EventToNextStateMap: EventToNextStateMap{
					EventNoOP: StateEnd,
				},
			},
		},
	}
}
