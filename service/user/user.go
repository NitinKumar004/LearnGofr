package user

import (
	userModel "GoFr/models/user"
	"errors"
	"gofr.dev/pkg/gofr"
)

type service struct {
	store Store
}

func New(s Store) *service {
	return &service{store: s}
}

func (s *service) InsertUser(ctx *gofr.Context, u userModel.User) (*userModel.User, error) {
	if u.Name == "" || u.Phone == "" || u.Email == "" {
		return nil, errors.New("all fields (name, phone, email) are required")
	}

	err := s.store.InsertUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return s.GetUserByID(ctx, u.ID)
}

func (s *service) GetUserByID(ctx *gofr.Context, id int) (*userModel.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return s.store.GetUserByID(ctx, id)
}

func (s *service) GetAllUsers(ctx *gofr.Context) ([]userModel.User, error) {
	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var filtered []userModel.User

	for _, t := range users {
		if t.Name != "" && t.Phone != "" && t.Email != "" {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func (s *service) DeleteAllUsers(ctx *gofr.Context) (string, error) {
	return s.store.DeleteAllUsers(ctx)
}

func (s *service) DeleteUserByID(ctx *gofr.Context, id int) (string, error) {
	if id <= 0 {
		return "", errors.New("invalid user ID")
	}

	return s.store.DeleteUserByID(ctx, id)
}
