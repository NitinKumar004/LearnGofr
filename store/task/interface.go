// store/task/store.go

package task

import (
	Task_Model "GoFr/models/task"
	"gofr.dev/pkg/gofr"
)

type TaskStore interface {
	Insertask(ctx *gofr.Context, t Task_Model.Task) error
	Gettaskbyid(ctx *gofr.Context, id int) (*Task_Model.Task, error)
	Getalltask(ctx *gofr.Context) ([]Task_Model.Task, error)
	Deletetask(ctx *gofr.Context, id int) error
	Completetask(ctx *gofr.Context, id int) error
}
