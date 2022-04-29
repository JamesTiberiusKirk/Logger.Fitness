package exercise_type

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_Controller(t *testing.T) {
	t.Run("testing_NewExerciseType", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDB := NewMockDatabaseInterface(ctrl)

	})

}
