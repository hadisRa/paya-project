package handler

import "paya/service"

type Handler struct {
	TaskHandler *TaskHandler
	UserHandler  *UserHandler
}

func NewHandlers(taskSrv service.Task, userSrv service.User) *Handler {
	return &Handler{
		TaskHandler: NewTaskHandler(taskSrv),
		UserHandler:  NewUserHandler(userSrv),
	}
}
