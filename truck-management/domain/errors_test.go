package domain

import (
	"reflect"
	"testing"
)

func TestNewConflict(t *testing.T) {
	type args struct {
		errorMessage string
	}
	tests := []struct {
		name string
		args args
		want *DomainErrors
	}{{
		"should return a domain error with 409 code",
		args{
			errorMessage: "error",
		},
		&DomainErrors{
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

func TestDomainErrors_Error(t *testing.T) {
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
				Code:         200,
			},
			want: "message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DomainErrors{
				ErrorMessage: tt.fields.ErrorMessage,
				Code:         tt.fields.Code,
			}
			if got := d.Error(); got != tt.want {
				t.Errorf("DomainErrors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
