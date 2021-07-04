package infrastructure

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

func Test_getError(t *testing.T) {
	GenericError := errors.New("generic error")

	type args struct {
		result *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"should return a Conflict error",
			args{
				result: &gorm.DB{
					Error: errors.New("Duplicate entry"),
				},
			},
			ConflictError,
		},
		{
			"should return generic error",
			args{
				result: &gorm.DB{
					Error: GenericError,
				},
			},
			GenericError,
		},
		{
			"should return Not Found error",
			args{
				result: &gorm.DB{
					RowsAffected: 0,
				},
			},
			NotFoundError,
		},
		{
			"should return nil",
			args{
				result: &gorm.DB{
					RowsAffected: 1,
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getError(tt.args.result); err != tt.want {
				t.Errorf("getError() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
