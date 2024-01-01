package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) addNewWorkoutLog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := 0
	exerciseName := "test exercise"
	CurrentWeight := 1000
	MaxReps := 100
	Notes := "test notes"

	result, err := app.workoutLogs.Insert(userId, exerciseName, CurrentWeight, MaxReps, Notes)
	if err != nil {
		app.serverError(w, err)
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		app.serverError(w, err)
	}
	w.Write(resultJson)
}
