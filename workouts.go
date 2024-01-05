package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models"
)

func (app *application) addNewWorkoutLog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

	errors := make(map[string]string)

	userId := workoutLog.UserId
	if userId < 0 {
		errors["user id"] = "invalid user id"
	}
	exerciseName := workoutLog.ExerciseName
	if strings.TrimSpace(exerciseName) == "" {
		errors["exercise name"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(exerciseName) > 100 {
		errors["exercise name"] = "This field is too long (maximum is 100 characters)"
	}
	CurrentWeight := workoutLog.CurrentWeight
	if CurrentWeight < 0 {
		errors["current weight"] = "invalid current weight"
	}
	MaxReps := workoutLog.MaxReps
	if MaxReps < 0 {
		errors["max reps"] = "invalid max reps"
	}
	Notes := workoutLog.Notes
	if utf8.RuneCountInString(Notes) > 250 {
		errors["notes"] = "This field is too long (maximum is 250 characters)"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

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
		w.Header().Set("Allow", "GET")
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
