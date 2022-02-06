export type ExerciseType = {
    exercise_type_id: string,
    // user_id: string,
    name: string,
    description: string,
    data_type: string,
    measurement_type: string,
}

export type ExerciseTypeMap = {
    [id: string]: ExerciseType
}