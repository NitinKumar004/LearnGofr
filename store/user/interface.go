package user

import models_user "GoFr/models/user"

type UserStore interface {
	InsertUser(u models_user.User) error
	GetUserByID(id int) (*models_user.User, error)
	GetAllUsers() ([]models_user.User, error)
	DeleteAllUsers() (string, error)
	DeleteUserByID(id int) (string, error)
}
