package api

import (
	"errors"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want ErrorResponse
	}{
		{
			"should return the same code and message when receiving a ClientError",
			args{
				err: &ClientErrors{
					ErrorMessage: "client error",
					Code:         400,
				},
			},
			ErrorResponse{
				"client error",
				400,
			},
		},
		{
			"should return 500 and internal server error when receiving a generic error",
			args{
				err: errors.New("generic"),
			},
			ErrorResponse{
				"internal server error",
				500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewErrorResponse(tt.args.err)
			if got != tt.want {
				t.Errorf("NewErrorResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
