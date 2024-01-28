package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models"
)

// TO-DO UPDATE this to match sign up repsonse.
type InvalidWorkoutLogDTO struct {
	FormData   WorkoutLogFormData
	FormErrors map[string]string
}

type WorkoutLogFormData struct {
	UserId        int
	ExerciseName  string
	CurrentWeight int
	MaxReps       int
	Notes         string
}

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
		errors["UserId"] = "invalid user id"
	}
	exerciseName := workoutLog.ExerciseName
	if strings.TrimSpace(exerciseName) == "" {
		errors["ExerciseName"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(exerciseName) > 100 {
		errors["ExerciseName"] = "This field is too long (maximum is 100 characters)"
	}
	CurrentWeight := workoutLog.CurrentWeight
	if CurrentWeight < 0 {
		errors["CurrentWeight"] = "invalid current weight"
	}
	MaxReps := workoutLog.MaxReps
	if MaxReps < 0 {
		errors["MaxReps"] = "invalid max reps"
	}
	Notes := workoutLog.Notes
	if utf8.RuneCountInString(Notes) > 250 {
		errors["Notes"] = "This field is too long (maximum is 250 characters)"
	}

	if len(errors) > 0 {
		result := InvalidWorkoutLogDTO{
			FormData: WorkoutLogFormData{
				UserId:        userId,
				ExerciseName:  exerciseName,
				CurrentWeight: CurrentWeight,
				MaxReps:       MaxReps,
				Notes:         Notes,
			},
			FormErrors: errors,
		}
		resultJson, err := json.Marshal(result)
		if err != nil {
			app.serverError(w, err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resultJson)
		return
	}

	_, err = app.workoutLogs.Insert(userId, exerciseName, CurrentWeight, MaxReps, Notes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var resultJson []byte
	resultJson, err = json.Marshal("Workoutlog successfully added!")
	if err != nil {
		app.serverError(w, err)
	}

	//app.session.Put(r, "flash", "Workoutlog successfully created!")

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
