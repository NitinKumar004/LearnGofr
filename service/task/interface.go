package task

import (
	taskModel "GoFr/models/task"
	"gofr.dev/pkg/gofr"
)

// create a interface of store here that we rhave implement in this code
// store interface defines all the methods the service layer depends on.
// This allows the service to interact with any implementation (real DB, mock, etc.),
// as long as it satisfies this contract.
// It helps achieve decoupling and testability.

type Store interface {
	Insertask(ctx *gofr.Context, t taskModel.Task) error
	Getalltask(ctx *gofr.Context) ([]taskModel.Task, error)
	Gettaskbyid(ctx *gofr.Context, id int) (*taskModel.Task, error)
	Deletetask(ctx *gofr.Context, id int) (string, error)
	Completetask(ctx *gofr.Context, id int) error
}
