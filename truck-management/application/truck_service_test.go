package application

import (
	"errors"
	"reflect"
	"testing"
	"time"
	"truck-management/truck-management/domain"
)

var timeNow = time.Now().Format(time.RFC3339)

func TestTruckService_CreateNewTruck(t *testing.T) {
	type fields struct {
		TruckRepository ITruckRepository
		TruckValidator  ITruckValidator
	}
	type args struct {
		newTruck domain.Truck
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Truck
		wantErr bool
	}{
		{
			name: "should return an error when Validator returns any error",
			fields: fields{
				TruckValidator: TruckValidatorMock{
					IsValidTruckMock: func(newTruck domain.Truck) error {
						return errors.New("license plate is already in use")
					},
				},
			},
			args: args{
				newTruck: domain.Truck{},
			},
			want:    domain.Truck{},
			wantErr: true,
		},
		{
			name: "should return an error when CreateTruck returns any error",
			fields: fields{
				TruckRepository: &TruckRepositoryMock{
					CreateTruckMock: func(*domain.Truck) error {
						return errors.New("failed to create truck")
					},
				},
				TruckValidator: TruckValidatorMock{
					IsValidTruckMock: func(newTruck domain.Truck) error {
						return nil
					},
				},
			},
			args: args{
				newTruck: domain.Truck{},
			},
			want:    domain.Truck{},
			wantErr: true,
		},
		{
			name: "should return TruckEntity when creating with success",
			fields: fields{
				TruckRepository: &TruckRepositoryMock{
					CreateTruckMock: func(truck *domain.Truck) error {
						truck.ID = "f4bd8a69-372f-4d4b-92ca-5c2424a3a539"
						truck.CreatedAt = timeNow
						truck.UpdatedAt = timeNow
						return nil
					},
				},
				TruckValidator: TruckValidatorMock{
					IsValidTruckMock: func(newTruck domain.Truck) error {
						return nil
					},
				},
			},
			args: args{
				newTruck: domain.Truck{
					LicensePlate: "ABC1234",
					EldID:        "00001234",
					Carrier:      "Third Carrier",
					Size:         23,
					Color:        "blue",
					Make:         "any-make",
					Model:        "any-model",
					Year:         2020,
				},
			},
			want: domain.Truck{
				ID:           "f4bd8a69-372f-4d4b-92ca-5c2424a3a539",
				LicensePlate: "ABC1234",
				EldID:        "00001234",
				Carrier:      "Third Carrier",
				Size:         23,
				Color:        "blue",
				Make:         "any-make",
				Model:        "any-model",
				Year:         2020,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TruckService{
				TruckRepository: tt.fields.TruckRepository,
				TruckValidator:  tt.fields.TruckValidator,
			}
			got, err := c.CreateNewTruck(tt.args.newTruck)
			if (err != nil) != tt.wantErr {
				t.Errorf("TruckService.CreateNewTruck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruckService.CreateNewTruck() = %v, want %v", got, tt.want)
			}
		})
	}
}

type GetTruckByLicensePlateAndEldIDMock func(licensePlate, eldID string) (domain.Truck, error)

type TruckRepositoryMock struct {
	CreateTruckMock func(*domain.Truck) error
}

func (t TruckRepositoryMock) CreateTruck(truck *domain.Truck) error {
	return t.CreateTruckMock(truck)
}

type TruckValidatorMock struct {
	IsValidTruckMock func(newTruck domain.Truck) error
}

func (t TruckValidatorMock) IsValidTruck(newTruck domain.Truck) error {
	return t.IsValidTruckMock(newTruck)
}
