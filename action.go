package main

import "fmt"

type Action interface {
	Execute(eventCtx EventContext) EventType
}

type ActionOrderCreated struct{}

func (a *ActionOrderCreated) Execute(eventCtx EventContext) EventType {
	order := eventCtx.(*OrderCreateContext)

	if len(order.items) == 0 {
		order.err = fmt.Errorf("no items purchased")
		return EventOrderFailed
	}

	return EventOrderPlaced
}

type ActionOrderFailed struct{}

func (a *ActionOrderFailed) Execute(eventCtx EventContext) EventType {
	return EventNoOP
}

type ActionOrderPlaced struct{}

func (a *ActionOrderPlaced) Execute(eventCtx EventContext) EventType {
	return EventChargeCard
}

type ActionCharingCard struct{}

func (a *ActionCharingCard) Execute(eventCtx EventContext) EventType {
	order := eventCtx.(*OrderCreateContext)

	if len(order.card) != 16 {
		order.err = fmt.Errorf("invalid card details")
		return EventTransactionFailed
	}

	return EventOrderShipped
}

type ActionTransactionFailed struct{}

func (a *ActionTransactionFailed) Execute(eventCtx EventContext) EventType {
	return EventNoOP
}

type ActionOrderShipped struct{}

func (a *ActionOrderShipped) Execute(eventCtx EventContext) EventType {
	return EventNoOP
}
