package application

import (
	"errors"
	"reflect"
	"testing"
	"truck-management/truck-management/domain"
)

func TestCreateTruckService_CreateNewTruck(t *testing.T) {
	type fields struct {
		TruckRepository TruckRepository
		TruckValidator  TruckValidator
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
			name: "should return an error when Get returns any error",
			fields: fields{
				TruckRepository: &TruckRepositoryMock{
					GetTruckByLicensePlateAndEldIDMock: func(licensePlate, eldID string) (domain.Truck, error) {
						return domain.Truck{}, errors.New("something went wrong")
					},
				},
				TruckValidator: func(newTruck domain.Truck, existingTruck domain.Truck) error {
					return nil
				},
			},
			args: args{
				newTruck: domain.Truck{},
			},
			want:    domain.Truck{},
			wantErr: true,
		},
		{
			name: "should return an error when Validator returns any error",
			fields: fields{
				TruckRepository: &TruckRepositoryMock{
					GetTruckByLicensePlateAndEldIDMock: func(licensePlate, eldID string) (domain.Truck, error) {
						return domain.Truck{}, nil
					},
				},
				TruckValidator: func(newTruck domain.Truck, existingTruck domain.Truck) error {
					return errors.New("license plate is already in use")
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
					GetTruckByLicensePlateAndEldIDMock: func(licensePlate, eldID string) (domain.Truck, error) {
						return domain.Truck{}, nil
					},
					CreateTruckMock: func(domain.Truck) error {
						return errors.New("failed to create truck")
					},
				},
				TruckValidator: func(newTruck domain.Truck, existingTruck domain.Truck) error {
					return nil
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
					GetTruckByLicensePlateAndEldIDMock: func(licensePlate, eldID string) (domain.Truck, error) {
						return domain.Truck{}, nil
					},
					CreateTruckMock: func(domain.Truck) error {
						return nil
					},
				},
				TruckValidator: func(newTruck domain.Truck, existingTruck domain.Truck) error {
					return nil
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
				LicensePlate: "ABC1234",
				EldID:        "00001234",
				Carrier:      "Third Carrier",
				Size:         23,
				Color:        "blue",
				Make:         "any-make",
				Model:        "any-model",
				Year:         2020,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateTruckService{
				TruckRepository: tt.fields.TruckRepository,
				TruckValidator:  tt.fields.TruckValidator,
			}
			got, err := c.CreateNewTruck(tt.args.newTruck)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTruckService.CreateNewTruck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTruckService.CreateNewTruck() = %v, want %v", got, tt.want)
			}
		})
	}
}

type GetTruckByLicensePlateAndEldIDMock func(licensePlate, eldID string) (domain.Truck, error)

type TruckRepositoryMock struct {
	GetTruckByLicensePlateAndEldIDMock func(licensePlate, eldID string) (domain.Truck, error)
	CreateTruckMock                    func(domain.Truck) error
}

func (t TruckRepositoryMock) GetTruckByLicensePlateAndEldID(licensePlate, eldID string) (domain.Truck, error) {
	return t.GetTruckByLicensePlateAndEldIDMock(licensePlate, eldID)
}

func (t TruckRepositoryMock) CreateTruck(truck domain.Truck) error {
	return t.CreateTruckMock(truck)
}
