package main

type StateType string

// nolint
const (
	StateInitial           StateType = "StateInitial"
	StateOrderCreated      StateType = "StateOrderCreated"
	StateOrderFailed       StateType = "StateOrderFailed"
	StateOrderPlaced       StateType = "StateOrderPlaced"
	StateChargingCard      StateType = "StateChargingCard"
	StateTransactionFailed StateType = "StateTransactionFailed"
	StateOrderShipped      StateType = "StateOrderShipped"
	StateError             StateType = "StateError"
	StateEnd               StateType = "StateEnd"
)
