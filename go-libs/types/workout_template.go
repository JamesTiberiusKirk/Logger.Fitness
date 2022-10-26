package types

type WorkoutTemplate struct {
	ID        string
	Name      string
	Exercises [][]ExerciseType
	Notes     string
}
