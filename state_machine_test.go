package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidWorkflow(t *testing.T) {
	stateMachine := NewStateMachine()

	order := OrderCreateContext{
		items: []string{"a", "b", "c"},
		card:  "1111222233334444",
	}

	stateTransitions := []StateType{}
	expectedStateTransitions := []StateType{
		StateInitial,
		StateOrderCreated,
		StateOrderPlaced,
		StateChargingCard,
		StateOrderShipped,
	}

	err := stateMachine.SendEvent(EventOrderCreated, &order, &stateTransitions)
	assert.Equal(t, stateTransitions, expectedStateTransitions)
	assert.Nil(t, err)
}

func TestOrderFailedWorkflow(t *testing.T) {
	stateMachine := NewStateMachine()

	order := OrderCreateContext{}

	stateTransitions := []StateType{}
	expectedStateTransitions := []StateType{
		StateInitial,
		StateOrderCreated,
		StateOrderFailed,
	}

	err := stateMachine.SendEvent(EventOrderCreated, &order, &stateTransitions)
	assert.Equal(t, expectedStateTransitions, stateTransitions)
	assert.Nil(t, err)
}

func TestFailedTransactions(t *testing.T) {
	stateMachine := NewStateMachine()

	order := OrderCreateContext{
		items: []string{"a", "b", "c"},
	}

	stateTransitions := []StateType{}
	expectedStateTransitions := []StateType{
		StateInitial,
		StateOrderCreated,
		StateOrderPlaced,
		StateChargingCard,
		StateTransactionFailed,
	}

	err := stateMachine.SendEvent(EventOrderCreated, &order, &stateTransitions)
	assert.Equal(t, stateTransitions, expectedStateTransitions)
	assert.Nil(t, err)

}
