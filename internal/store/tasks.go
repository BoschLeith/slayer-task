package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Task struct {
	ID        int64    `json:"id"`
	Equipment string   `json:"equipment"`
	Inventory string   `json:"inventory"`
	Monster   string   `json:"monster"`
	Notes     []string `json:"notes"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type TaskStore struct {
	db *sql.DB
}

func (s *TaskStore) Create(ctx context.Context, task *Task) error {
	query := `
		INSERT INTO tasks (equipment, inventory, monster, notes)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		task.Equipment,
		task.Inventory,
		task.Monster,
		pq.Array(task.Notes),
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskStore) GetByID(ctx context.Context, id int64) (*Task, error) {
	query := `SELECT id, equipment, inventory, monster, notes, created_at, updated_at
	FROM TASKS
	WHERE id = $1
	`

	var task Task
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Equipment,
		&task.Inventory,
		&task.Monster,
		pq.Array(&task.Notes),
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &task, nil
}
