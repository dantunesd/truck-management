package application

import (
	"errors"
	"reflect"
	"testing"
	"truck-management/truck-management/domain"
)

func TestTripService_GetTrip(t *testing.T) {
	type fields struct {
		tripRepository ITripRepository
		truckService   ITruckService
	}
	type args struct {
		truckID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Trip
		wantErr bool
	}{
		{
			name: "should return an error when GetTruck returns any error",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (*domain.Truck, error) {
						return &domain.Truck{}, errors.New("truck not found")
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (*domain.Trip, error) {
						return &domain.Trip{}, nil
					},
				},
			},
			args: args{
				truckID: 1,
			},
			want:    &domain.Trip{},
			wantErr: true,
		},
		{
			name: "should return an error when GetTrip returns any error",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (*domain.Truck, error) {
						return &domain.Truck{}, nil
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (*domain.Trip, error) {
						return &domain.Trip{}, errors.New("not found")
					},
				},
			},
			args: args{
				truckID: 1,
			},
			want:    &domain.Trip{},
			wantErr: true,
		},
		{
			name: "should return an TripEntity when getting with success",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (*domain.Truck, error) {
						return &domain.Truck{}, nil
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (*domain.Trip, error) {
						return &domain.Trip{
							ID:           1,
							TruckID:      1,
							Origin:       90,
							Destination:  90,
							State:        domain.ONGOING,
							Odometer:     100,
							EngineHours:  2,
							AverageSpeed: 150,
							UpdatedAt:    timeNow,
						}, nil
					},
				},
			},
			args: args{
				truckID: 1,
			},
			want: &domain.Trip{
				ID:           1,
				TruckID:      1,
				Origin:       90,
				Destination:  90,
				State:        domain.ONGOING,
				Odometer:     100,
				EngineHours:  2,
				AverageSpeed: 150,
				UpdatedAt:    timeNow,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTripService(tt.fields.tripRepository, tt.fields.truckService)
			got, err := tr.GetTrip(tt.args.truckID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TripService.GetTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TripService.GetTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTripService_UpdateTrip(t *testing.T) {
	type fields struct {
		tripRepository ITripRepository
		truckService   ITruckService
	}
	type args struct {
		location domain.Location
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should return nil",
			fields: fields{},
			args: args{
				location: domain.Location{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TripService{
				tripRepository: tt.fields.tripRepository,
				truckService:   tt.fields.truckService,
			}
			if err := tr.UpdateTrip(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("TripService.UpdateTrip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TripRepositoryMock struct {
	GetTripMock func(truckID int) (*domain.Trip, error)
}

func (t TripRepositoryMock) GetTrip(truckID int) (*domain.Trip, error) {
	return t.GetTripMock(truckID)
}
