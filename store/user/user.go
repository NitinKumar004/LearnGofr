package user

import (
	userModel "GoFr/models/user"
	"database/sql"
	"errors"
	"gofr.dev/pkg/gofr"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) InsertUser(ctx *gofr.Context, u userModel.User) error {
	//fmt.Println(u)
	_, err := s.db.Exec("INSERT INTO usermanage(userid, username, userphone, useremail) VALUES (?, ?, ?, ?)",
		u.ID, u.Name, u.Phone, u.Email)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByID(ctx *gofr.Context, id int) (*userModel.User, error) {
	var u userModel.User
	row := s.db.QueryRow("SELECT userid, username, userphone, useremail FROM usermanage WHERE userid = ?", id)

	err := row.Scan(&u.ID, &u.Name, &u.Phone, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Store) GetAllUsers(ctx *gofr.Context) ([]userModel.User, error) {
	var allUsers []userModel.User
	rows, err := s.db.Query("SELECT userid, username, userphone, useremail FROM usermanage")
	if err != nil {
		return nil, errors.New("error fetching users")
	}

	for rows.Next() {
		var u userModel.User
		err := rows.Scan(&u.ID, &u.Name, &u.Phone, &u.Email)
		if err != nil {
			return nil, errors.New("error scanning user data")
		}

		allUsers = append(allUsers, u)
	}
	return allUsers, nil
}

func (s *Store) DeleteAllUsers(ctx *gofr.Context) (string, error) {
	result, err := s.db.Exec("DELETE FROM usermanage")
	if err != nil {
		return "error deleting users", errors.New("failed to delete users")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check deletion result")
	}
	if rowsAffected == 0 {
		return "no users found", errors.New("no users to delete")
	}

	return "all users deleted successfully", nil
}
func (s *Store) DeleteUserByID(ctx *gofr.Context, id int) (string, error) {
	result, err := s.db.Exec("DELETE FROM usermanage WHERE userid = ?", id)
	if err != nil {
		return "error deleting user", errors.New("failed to delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check deletion result")
	}
	if rowsAffected == 0 {
		return "no user found", errors.New("no user with this ID exists")
	}

	return "user deleted successfully", nil
}
