package task

import (
	taskModel "GoFr/models/task"
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

func (s *Store) Insertask(ctx *gofr.Context, t taskModel.Task) error {
	_, err := s.db.Exec("INSERT INTO taskmanage(taskid, taskname, status, assigned_user_id) VALUES (?, ?, ?, ?)",
		t.ID, t.Name, t.Status, t.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Gettaskbyid(ctx *gofr.Context, id int) (*taskModel.Task, error) {
	var t taskModel.Task
	row := s.db.QueryRow("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage WHERE taskid=?", id)
	err := row.Scan(&t.ID, &t.Name, &t.Status, &t.UserID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Store) Getalltask(ctx *gofr.Context) ([]taskModel.Task, error) {
	rows, err := s.db.Query("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage")
	if err != nil {
		return nil, errors.New("error fetching all tasks")
	}
	defer rows.Close()

	var tasks []taskModel.Task
	for rows.Next() {
		var t taskModel.Task
		if err := rows.Scan(&t.ID, &t.Name, &t.Status, &t.UserID); err != nil {
			return nil, errors.New("error scanning tasks")
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (s *Store) Deletetask(ctx *gofr.Context, id int) (string, error) {
	result, err := s.db.Exec("DELETE FROM taskmanage WHERE taskid = ?", id)
	if err != nil {
		return "error to delete", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "row not affected", err
	}
	if rowsAffected == 0 {
		return "id not exists", errors.New("no task with this ID exists")
	}

	return "deleted successfully", nil
}

func (s *Store) Completetask(ctx *gofr.Context, id int) error {
	result, err := s.db.Exec("UPDATE taskmanage SET status = ? WHERE taskid = ?", "complete", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no task with this ID exists")
	}

	return nil
}
