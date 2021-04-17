package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stateMachine = NewStateMachine()
)

func TestCreateFailedWorkflow(t *testing.T) {
	invalidOrder := OrderCreateContext{}

	err := stateMachine.SendEvent(EventOrderCreated, &invalidOrder)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "No transition exists")

}
