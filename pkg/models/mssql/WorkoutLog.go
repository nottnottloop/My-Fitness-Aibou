package msssql

import (
	"database/sql"

	"github.com/Lionel-Wilson/My-Fitness-Aibou/pkg/models"
)

type WorkoutLogModel struct {
	DB *sql.DB
}

func (m *WorkoutLogModel) Insert(UserId int, ExerciseName string, CurrentWeight int, MaxReps int, Notes string) (int, error) {

	query := `
		INSERT INTO workoutlogs (user_id, exercise_name, current_weight, max_reps, notes, created)
		VALUES(?, ?, ?, ?, ?, UTC_TIMESTAMP());
		`

	result, err := m.DB.Exec(query, UserId, ExerciseName, CurrentWeight, MaxReps, Notes)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *WorkoutLogModel) Get(Id int) (*models.WorkoutLog, error) {
	result := &models.WorkoutLog{}

	query := `SELECT id,user_id, exercise_name, current_weight, max_reps, notes, created FROM workoutlogs WHERE id = ?;`

	err := m.DB.QueryRow(query, Id).Scan(&result.ID, &result.UserId, &result.ExerciseName, &result.CurrentWeight, &result.MaxReps, &result.Notes, &result.Created)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *WorkoutLogModel) GetAll() ([]*models.WorkoutLog, error) {
	return nil, nil
}
