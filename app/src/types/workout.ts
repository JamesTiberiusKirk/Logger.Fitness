import { Exercise } from "./exercise"

export type Workout = {
    workout_id: string,
    // user_id:string,
    start_time: string,
    end_time: string,
    notes: string,
    title: string,
}

export type WorkoutGroup = {
    workout: Workout,
    exercises:Exercise[],
}