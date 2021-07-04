package infrastructure

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

func Test_hasError(t *testing.T) {
	type args struct {
		result *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should return true when exists error",
			args{
				result: &gorm.DB{
					Error: errors.New("generic error"),
				},
			},
			true,
		},
		{
			"should return false when not exists error",
			args{
				result: &gorm.DB{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasError(tt.args.result); got != tt.want {
				t.Errorf("hasError() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			"should return true when exists error",
			args{
				result: &gorm.DB{
					Error: errors.New("Duplicate entry"),
				},
			},
			true,
		},
		{
			"should return false when not exists error",
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
			"should return true when is rows was not affected",
			args{
				result: &gorm.DB{
					RowsAffected: 0,
				},
			},
			true,
		},
		{
			"should return false when rows was affected",
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
