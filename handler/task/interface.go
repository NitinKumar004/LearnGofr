package task

import (
	taskModel "GoFr/models/task"
	"gofr.dev/pkg/gofr"
)

type service interface {
	Insertask(ctx *gofr.Context, t taskModel.Task) (*taskModel.Task, error)
	Getalltask(ctx *gofr.Context) ([]taskModel.Task, error)
	Gettaskbyid(ctx *gofr.Context, id int) (*taskModel.Task, error)
	Deletetask(ctx *gofr.Context, id int) (string, error)
	Completetask(ctx *gofr.Context, id int) (*taskModel.Task, error)
}
