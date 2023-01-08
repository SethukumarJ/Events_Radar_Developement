package handler

import (
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type EventHandler struct {
	eventUseCase usecase.EventUseCase
}

func NewEventHandler(usecase usecase.EventUseCase) EventHandler {
	return EventHandler{
		eventUseCase: usecase,
	}
}
