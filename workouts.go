package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models"
)

func (app *application) addNewWorkoutLog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var workoutLog models.WorkoutLog

	err = json.Unmarshal(body, &workoutLog)
	if err != nil {
		app.serverError(w, err)
		return
	}

	userId := workoutLog.UserId
	exerciseName := workoutLog.ExerciseName
	CurrentWeight := workoutLog.CurrentWeight
	MaxReps := workoutLog.MaxReps
	Notes := workoutLog.Notes

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

func (app *application) getWorkoutLog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "Invalid id provided", http.StatusBadRequest)
		return
	}

	result, err := app.workoutLogs.Get(id)
	if err == models.ErrNoRecord {
		http.NotFound(w, r)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(resultJson)

}

func (app *application) getAllWorkoutLogs(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := app.workoutLogs.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(resultJson)

}
