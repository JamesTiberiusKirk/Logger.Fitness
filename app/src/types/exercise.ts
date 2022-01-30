export type Exercise = {
    id: string,
    workout_id: string,
    exercise_type_id: string,
    sets: Set[],
    single_value: SingleValue,
    note: string,
}

export type Set = {
    resistance: number,
    reps: number,
    is_drop_set: boolean,
}

export type SingleValue = {
    value: number,
}
