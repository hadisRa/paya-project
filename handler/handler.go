package handler

import "paya/service"

type Handler struct {
	TasskHandler *TaskHandler
	UserHandler  *UserHandler
}

func NewHandlers(taskSrv service.Task, userSrv service.User) *Handler {
	return &Handler{
		TasskHandler: NewTaskHandler(taskSrv),
		UserHandler:  NewUserHandler(userSrv),
	}
}
