package task

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"strconv"

	taskModel "GoFr/models/task"
)

type Handler struct {
	service service
}

func New(s service) *Handler {
	return &Handler{service: s}
}
func (h *Handler) Addtask(ctx *gofr.Context) (interface{}, error) {
	var t taskModel.Task

	if err := ctx.Bind(&t); err != nil {
		return nil, err
	}

	task, err := h.service.Insertask(ctx, t)
	if err != nil {
		return nil, err
	}

	return task, nil
}
func (h *Handler) GetAllTask(ctx *gofr.Context) (interface{}, error) {
	tasks, err := h.service.Getalltask(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (h *Handler) GetTaskById(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	task, err := h.service.Gettaskbyid(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil

}

func (h *Handler) CompleteTask(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	task, err := h.service.Completetask(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil

}
func (h *Handler) DeleteTask(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	msg, err := h.service.Deletetask(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("handler", msg)
	return msg, nil

}
