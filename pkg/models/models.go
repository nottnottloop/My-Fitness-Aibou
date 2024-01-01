package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type WorkoutLog struct {
	ID            int
	UserId        int
	ExerciseName  string
	CurrentWeight int
	MaxReps       int
	Notes         string
	Created       time.Time
}
