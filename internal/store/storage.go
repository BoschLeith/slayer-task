package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		Create(context.Context, *Task) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TaskStore{db},
	}
}
