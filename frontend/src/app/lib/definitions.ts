export type WorkoutLogFormData = {
    user_id: number;
    exercise_name: string;
    current_weight: number;
    max_reps: number;
    notes: string;
  };
  
export type AddWorkoutLogResponse = {
    flashMessage: string
  };


  
