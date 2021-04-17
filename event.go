package main

type EventType string

// nolint
const (
	EventOrderCreated      EventType = "EventOrderCreated"
	EventOrderFailed       EventType = "EventOrderFailed"
	EventOrderPlaced       EventType = "EventOrderPlaced"
	EventChargeCard        EventType = "EventChargeCard"
	EventTransactionFailed EventType = "EventTransactionFailed"
	EventOrderShipped      EventType = "EventOrderShipped"
	EventNoOP              EventType = "EventNoOP"
)

type EventContext interface{}

type EventToNextStateMap map[EventType]StateType

type OrderCreateContext struct {
	items []string
	err   error
}
