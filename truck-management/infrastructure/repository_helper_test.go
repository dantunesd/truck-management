package infrastructure

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

func Test_isDuplicated(t *testing.T) {
	type args struct {
		result *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should return true when exists an error",
			args{
				result: &gorm.DB{
					Error: errors.New("Duplicate entry"),
				},
			},
			true,
		},
		{
			"should return false when doesnt exist an error",
			args{
				result: &gorm.DB{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDuplicated(tt.args.result); got != tt.want {
				t.Errorf("isDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNotFound(t *testing.T) {
	type args struct {
		result *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should return true when 0 rows was affected",
			args{
				result: &gorm.DB{
					RowsAffected: 0,
				},
			},
			true,
		},
		{
			"should return false when some rows was affected",
			args{
				result: &gorm.DB{
					RowsAffected: 1,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNotFound(tt.args.result); got != tt.want {
				t.Errorf("isNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
