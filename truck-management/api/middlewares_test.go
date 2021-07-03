package api

import (
	"errors"
	"testing"
)

func TestGetErrorResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			"should return code and message when receiving a client error",
			args{
				err: &ClientErrors{
					ErrorMessage: "client error",
					Code:         400,
				},
			},
			400,
			"client error",
		},
		{
			"should return code and message when receiving a generic error",
			args{
				err: errors.New("generic"),
			},
			500,
			"internal server error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetErrorResponse(tt.args.err)
			if got != tt.want {
				t.Errorf("GetErrorResponse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetErrorResponse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
