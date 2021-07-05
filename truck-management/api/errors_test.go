package api

import (
	"reflect"
	"testing"
)

func TestNewBadRequest(t *testing.T) {
	type args struct {
		errorMessage string
	}
	tests := []struct {
		name string
		args args
		want *ClientErrors
	}{
		{
			"should return a ClientError with 400 code",
			args{
				errorMessage: "error",
			},
			&ClientErrors{
				ErrorMessage: "error",
				Code:         400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadRequest(tt.args.errorMessage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConflict(t *testing.T) {
	type args struct {
		errorMessage string
	}
	tests := []struct {
		name string
		args args
		want *ClientErrors
	}{
		{
			"should return a ClientError with 409 code",
			args{
				errorMessage: "error",
			},
			&ClientErrors{
				ErrorMessage: "error",
				Code:         409,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConflict(tt.args.errorMessage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNotFound(t *testing.T) {
	type args struct {
		errorMessage string
	}
	tests := []struct {
		name string
		args args
		want *ClientErrors
	}{
		{
			"should return a ClientError with 404 code",
			args{
				errorMessage: "error",
			},
			&ClientErrors{
				ErrorMessage: "error",
				Code:         404,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotFound(tt.args.errorMessage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientErrors_Error(t *testing.T) {
	type fields struct {
		ErrorMessage string
		Code         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return the message",
			fields: fields{
				ErrorMessage: "message",
				Code:         400,
			},
			want: "message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ClientErrors{
				ErrorMessage: tt.fields.ErrorMessage,
				Code:         tt.fields.Code,
			}
			if got := d.Error(); got != tt.want {
				t.Errorf("ClientErrors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
