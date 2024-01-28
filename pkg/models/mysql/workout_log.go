package mysql

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

	query := `
	SELECT * FROM myfitnessaiboudb.workoutlogs;
	`

	rows, err := m.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workoutlogs := []*models.WorkoutLog{}

	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		w := &models.WorkoutLog{}
		// Use rows.Scan() to copy the values from each field in the row to the
		// new Snippet object that we created. Again, the arguments to row.Scan
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&w.ID, &w.UserId, &w.ExerciseName, &w.CurrentWeight, &w.MaxReps, &w.Notes, &w.Created)
		if err != nil {
			return nil, err
		}
		workoutlogs = append(workoutlogs, w)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workoutlogs, nil
}
