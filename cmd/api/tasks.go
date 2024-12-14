package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BoschLeith/slayer-task/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreateTaskPayload struct {
	Equipment string   `json:"equipment"`
	Inventory string   `json:"inventory"`
	Monster   string   `json:"monster"`
	Notes     []string `json:"notes"`
}

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTaskPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	task := &store.Task{
		Equipment: payload.Equipment,
		Inventory: payload.Inventory,
		Monster:   payload.Monster,
		Notes:     payload.Notes,
	}

	ctx := r.Context()

	if err := app.store.Tasks.Create(ctx, task); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, task); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "taskID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()

	task, err := app.store.Tasks.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, task); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
