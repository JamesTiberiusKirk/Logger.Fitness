package types

import "gorm.io/gorm"

// DataTypeLabel is the data type of the exercise type
type DataTypeLabel string

const (
	// SetLabel is for sets
	SetLabel DataTypeLabel = "sets"

	// CustomValueLabel for custom values. If this is not selected, then there will be no CustomDataFields
	CustomValueLabel DataTypeLabel = "custom-value"
)

// ExerciseType is for custom types of exercises
type ExerciseType struct {
	gorm.Model
	ID                  string                       `json:"id" gorm:"primaryKey"`
	Name                string                       `json:"name"`
	Description         string                       `json:"description"`
	DataType            DataTypeLabel                `json:"data_type"`
	CustomDataTypeShape map[string]CustomValueFields `json:"custom_data_type_shape"` // Only populated if DataType = "custom-value"
	MeasurementType     string                       `json:"measurement_type"`
	Tags                []string                     `json:"tags"`
}
