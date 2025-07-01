package user

import (
	"gofr.dev/pkg/gofr"
	"strconv"

	userModel "GoFr/models/user"
)

type Handler struct {
	service service
}

func New(s service) *Handler {
	return &Handler{service: s}
}

// POST /user
func (h *Handler) AddUser(ctx *gofr.Context) (interface{}, error) {
	var u userModel.User

	if err := ctx.Bind(&u); err != nil {
		return nil, err
	}

	// Call service
	user, err := h.service.InsertUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *Handler) GetUserByID(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *Handler) GetAllUsers(ctx *gofr.Context) (interface{}, error) {
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (h *Handler) DeleteAllUsers(ctx *gofr.Context) (interface{}, error) {
	msg, err := h.service.DeleteAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": msg}, nil
}

// DELETE /user/{id}
func (h *Handler) DeleteUserByID(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	msg, err := h.service.DeleteUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": msg}, nil
}
