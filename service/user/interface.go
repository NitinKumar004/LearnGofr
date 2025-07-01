package user

import (
	userModel "GoFr/models/user"
	"gofr.dev/pkg/gofr"
)

type Store interface {
	InsertUser(ctx *gofr.Context, u userModel.User) error
	GetUserByID(ctx *gofr.Context, id int) (*userModel.User, error)
	GetAllUsers(ctx *gofr.Context) ([]userModel.User, error)
	DeleteAllUsers(ctx *gofr.Context) (string, error)
	DeleteUserByID(ctx *gofr.Context, id int) (string, error)
}
