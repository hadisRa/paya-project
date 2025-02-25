package handler

import "paya/service"

type Handler struct {
	TasskHandler *TaskHandler
}

func NewHandlers(srv service.Task) *Handler {
	return &Handler{
		TasskHandler: NewTaskHandler(srv),
	}
}
