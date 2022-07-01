package common

import "Logger.Fitness/go-libs/types"

// CompareExerciseTypes compares two ExerciseType structs, returns true if they are the same. ID is ignored
func CompareExerciseTypes(a, b types.ExerciseType) bool {

	if a.DataType != b.DataType {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	if a.MeasurementType != b.MeasurementType {
		return false
	}
	if a.Name != b.Name {
		return false
	}

	return true
}
