package main

import (
	"net/http"

	"github.com/BoschLeith/slayer-task/internal/store"
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
