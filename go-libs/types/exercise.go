package types

import "gorm.io/gorm"

// TODO: figure out how to store sets and custom fields

// CustomValueFields fields for custom value
type CustomValueFields struct {
	DisplayName string  `json:"display_name"`
	DataType    string  `json:"data_type"`
	Value       float32 `json:"value"`
}

// Set for the data in exercises
type Set struct {
	Resps      int     `json:"resps"`
	Resistance float32 `json:"resistance"`
	IsDropSet  bool    `json:"is_drop_set"`
}

// Exercise is for one individual exercise in a workout
type Exercise struct {
	gorm.Model
	ID              string                       `json:"id" gorm:"primaryKey"`
	Meta            ExerciseType                 `json:"meta"`
	ExerciseTypeID  string                       `json:"exercise_type_id" gorm:"foreignKey"`
	OrderInWorkout  int                          `json:"order_in_workout"`
	OrderInSuperSet int                          `json:"order_in_super_set"`
	Sets            []Set                        `json:"sets"`
	CustomValue     map[string]CustomValueFields `json:"custom_value"`
	Notes           string                       `json:"notes"`
	Tags            []string                     `json:"tags"`
}
