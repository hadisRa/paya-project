package handler

import (
	"paya/repository"
	"paya/service"
)

type Handler struct {
	TaskHandler *TaskHandler
	UserHandler *UserHandler
}

func NewHandlers(taskSrv service.Task, userSrv service.User, cacheRepo repository.CacheInterface, userRepo repository.UserInterface) *Handler {
	return &Handler{
		TaskHandler: NewTaskHandler(taskSrv),
		UserHandler: NewUserHandler(userSrv, cacheRepo, userRepo),
	}
}
