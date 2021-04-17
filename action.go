package main

import "fmt"

type Action interface {
	Execute(eventCtx EventContext) EventType
}

type ActionCreateOrder struct{}

func (a *ActionCreateOrder) Execute(eventCtx EventContext) EventType {
	order := eventCtx.(*OrderCreateContext)

	fmt.Println("Validating order", order)
	if len(order.items) == 0 {
		order.err = fmt.Errorf("no items purchased")
		return EventOrderFailed
	}

	return EventOrderPlaced
}
