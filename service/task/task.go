package task

import (
	taskModel "GoFr/models/task"
	"errors"
	"fmt"
	"gofr.dev/pkg/gofr"
)

// create a interface of store here that we rhave implement in this code
// store interface defines all the methods the service layer depends on.
// This allows the service to interact with any implementation (real DB, mock, etc.),
// as long as it satisfies this contract.
// It helps achieve decoupling and testability.

// Service struct is the business logic layer.
// It holds a reference to a Store interface, so it can call DB-related methods indirectly.
// This ensures that the Service doesn’t depend on any specific DB implementation.
type Service struct {
	store Store
}

// New is a constructor function for the Service.
// It takes a Store implementation (could be real DB or a mock for testing)
// and returns a pointer to a new Service instance.
func New(s Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) Insertask(ctx *gofr.Context, t taskModel.Task) (*taskModel.Task, error) {
	if t.Name == "" || t.Status == "" {
		return nil, errors.New("task name and status cannot be empty")
	}
	err := s.store.Insertask(ctx, t)
	if err != nil {
		return nil, err
	}

	return s.Gettaskbyid(ctx, t.ID)
}
func (s *Service) Getalltask(ctx *gofr.Context) ([]taskModel.Task, error) {
	tasks, err := s.store.Getalltask(ctx)
	if err != nil {
		return nil, err
	}

	var filtered []taskModel.Task

	for _, t := range tasks {
		if t.Name != "" {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil

}
func (s *Service) Gettaskbyid(ctx *gofr.Context, id int) (*taskModel.Task, error) {
	if id <= 0 {
		return nil, errors.New("id should not be negative")
	}

	return s.store.Gettaskbyid(ctx, id)

}
func (s *Service) Deletetask(ctx *gofr.Context, id int) (string, error) {
	if id <= 0 {

		return "", errors.New("id should not be negative")
	}
	msg, err := s.store.Deletetask(ctx, id)

	if err != nil {
		return "error in this", err

	}
	fmt.Println("service", msg)
	return msg, nil

}
func (s *Service) Completetask(ctx *gofr.Context, id int) (*taskModel.Task, error) {

	if id <= 0 {
		return nil, errors.New("id should not be negative")
	}

	err := s.store.Completetask(ctx, id)
	if err != nil {
		return nil, err

	}
	return s.Gettaskbyid(ctx, id)

}
