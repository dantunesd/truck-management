package application

import (
	"errors"
	"reflect"
	"testing"
	"time"
	"truck-management/truck-management/domain"
)

var timeNow = time.Now().Format(time.RFC3339)

func TestTruckService_CreateTruck(t *testing.T) {
	type fields struct {
		truckRepository ITruckRepository
	}
	type args struct {
		truck domain.Truck
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Truck
		wantErr bool
	}{
		{
			name: "should return an error when CreateTruck returns any error",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					CreateTruckMock: func(*domain.Truck) error {
						return errors.New("failed to create truck")
					},
				},
			},
			args: args{
				truck: domain.Truck{},
			},
			want:    domain.Truck{},
			wantErr: true,
		},
		{
			name: "should return TruckEntity when creating with success",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					CreateTruckMock: func(truck *domain.Truck) error {
						truck.ID = 1
						truck.CreatedAt = timeNow
						truck.UpdatedAt = timeNow
						return nil
					},
				},
			},
			args: args{
				truck: domain.Truck{
					LicensePlate: "ABC1234",
					EldID:        "00001234",
					CarrierID:    "CARRIER-ID",
					Size:         23,
					Color:        "blue",
					Make:         "any-make",
					Model:        "any-model",
					Year:         2020,
				},
			},
			want: domain.Truck{
				ID:           1,
				LicensePlate: "ABC1234",
				EldID:        "00001234",
				CarrierID:    "CARRIER-ID",
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
			c := NewTruckService(tt.fields.truckRepository)
			got, err := c.CreateTruck(tt.args.truck)
			if (err != nil) != tt.wantErr {
				t.Errorf("TruckService.CreateTruck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruckService.CreateTruck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruckService_GetTruck(t *testing.T) {
	type fields struct {
		truckRepository ITruckRepository
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Truck
		wantErr bool
	}{
		{
			name: "should return an error when GetTruck returns any error",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					GetTruckMock: func(ID int) (domain.Truck, error) {
						return domain.Truck{}, errors.New("failed to get truck")
					},
				},
			},
			args: args{
				ID: 1,
			},
			want:    domain.Truck{},
			wantErr: true,
		},
		{
			name: "should return a Truck getting with success",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					GetTruckMock: func(ID int) (domain.Truck, error) {
						return domain.Truck{
							ID:           1,
							LicensePlate: "ABC12345",
						}, nil
					},
				},
			},
			args: args{
				ID: 1,
			},
			want: domain.Truck{
				ID:           1,
				LicensePlate: "ABC12345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TruckService{
				truckRepository: tt.fields.truckRepository,
			}
			got, err := c.GetTruck(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TruckService.GetTruck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruckService.GetTruck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruckService_DeleteTruck(t *testing.T) {
	type fields struct {
		truckRepository ITruckRepository
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return an error when DeleteTruck returns any error",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					DeleteTruckMock: func(ID int) error {
						return errors.New("failed to delete truck")
					},
				},
			},
			args: args{
				ID: 1,
			},
			wantErr: true,
		},
		{
			name: "should return nil when deleting with success",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					DeleteTruckMock: func(ID int) error {
						return nil
					},
				},
			},
			args: args{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TruckService{
				truckRepository: tt.fields.truckRepository,
			}
			if err := c.DeleteTruck(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("TruckService.DeleteTruck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTruckService_UpdateTruck(t *testing.T) {
	type fields struct {
		truckRepository ITruckRepository
	}
	type args struct {
		ID    int
		truck domain.Truck
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return an error when UpdateTruck returns any error",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					UpdateTruckMock: func(ID int, truck *domain.Truck) error {
						return errors.New("failed to create truck")
					},
				},
			},
			args: args{
				ID:    1,
				truck: domain.Truck{},
			},
			wantErr: true,
		},
		{
			name: "should return nil when updating with success",
			fields: fields{
				truckRepository: &TruckRepositoryMock{
					UpdateTruckMock: func(ID int, truck *domain.Truck) error {
						return nil
					},
				},
			},
			args: args{
				ID:    1,
				truck: domain.Truck{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TruckService{
				truckRepository: tt.fields.truckRepository,
			}
			if err := c.UpdateTruck(tt.args.ID, tt.args.truck); (err != nil) != tt.wantErr {
				t.Errorf("TruckService.UpdateTruck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TruckRepositoryMock struct {
	CreateTruckMock func(*domain.Truck) error
	GetTruckMock    func(ID int) (domain.Truck, error)
	DeleteTruckMock func(ID int) error
	UpdateTruckMock func(ID int, truck *domain.Truck) error
}

func (t TruckRepositoryMock) CreateTruck(truck *domain.Truck) error {
	return t.CreateTruckMock(truck)
}

func (t TruckRepositoryMock) GetTruck(ID int) (domain.Truck, error) {
	return t.GetTruckMock(ID)
}

func (t TruckRepositoryMock) DeleteTruck(ID int) error {
	return t.DeleteTruckMock(ID)
}

func (t TruckRepositoryMock) UpdateTruck(ID int, truck *domain.Truck) error {
	return t.UpdateTruckMock(ID, truck)
}
