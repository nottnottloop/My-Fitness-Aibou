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

	/*
		rows, err := m.DB.Query("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE'")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Iterate through the result set and print table names
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Table:", tableName)
		}

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		return 1, nil
	*/
}

func (m *WorkoutLogModel) Get(Id int) (*models.WorkoutLog, error) {
	return nil, nil
}

func (m *WorkoutLogModel) GetAll() ([]*models.WorkoutLog, error) {
	return nil, nil
}
