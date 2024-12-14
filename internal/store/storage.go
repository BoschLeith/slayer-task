package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Tasks interface {
		GetByID(context.Context, int64) (*Task, error)
		Create(context.Context, *Task) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TaskStore{db},
	}
}
