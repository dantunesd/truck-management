package api

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTruckIdHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should not abort with error when ID is valid",
			args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{
							Key:   "id",
							Value: "1",
						},
					},
				},
			},
			false,
		},
		{
			"should abort with error when ID is not present",
			args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{},
					},
				},
			},
			true,
		},

		{
			"should abort with error when ID is invalid",
			args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{
							Key:   "id",
							Value: "invalid",
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TruckIdHandler(tt.args.c)
			gotErr := tt.args.c.Errors.Last() != nil
			if gotErr != tt.wantErr {
				t.Errorf("TruckIdHandler() got = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	ginContextWithError, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContextWithError.Error(errors.New("error"))

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			"should response 500 when exists a generic error in context",
			args{
				c: ginContextWithError,
			},
			500,
		},
		{
			"should response 500 when doesn't exist a generic error in context",
			args{
				c: ginContext,
			},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorHandler(tt.args.c)

			got := tt.args.c.Writer.Status()
			if got != tt.wantStatus {
				t.Errorf("ErrorHandler() got = %v, want %v", got, tt.wantStatus)
			}
		})
	}
}

func TestLogHandler(t *testing.T) {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	ginContextWithError, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContextWithError.Error(errors.New("error"))

	type fields struct {
		logger ILogger
	}

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"should log when exists an error in context",
			fields{
				logger: LoggerMock{},
			},
			args{
				c: ginContextWithError,
			},
		},
		{
			"should not log when doesn't exist an error in context",
			fields{
				logger: LoggerMock{},
			},
			args{
				c: ginContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := LogHandler(tt.fields.logger)
			h(tt.args.c)
		})
	}
}

type LoggerMock struct{}

func (l LoggerMock) Error(args ...interface{}) {}
