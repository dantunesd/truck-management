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
		TripUpdater    ITripUpdater
	}
	type args struct {
		truckID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Trip
		wantErr bool
	}{
		{
			name: "should return an error when GetTruck fails",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (domain.Truck, error) {
						return domain.Truck{}, errors.New("truck not found")
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{}, nil
					},
				},
			},
			args: args{
				truckID: 1,
			},
			want:    domain.Trip{},
			wantErr: true,
		},
		{
			name: "should return an error when GetTrip fails",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (domain.Truck, error) {
						return domain.Truck{}, nil
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{}, errors.New("generic error")
					},
				},
			},
			args: args{
				truckID: 1,
			},
			want:    domain.Trip{},
			wantErr: true,
		},
		{
			name: "should return an TripEntity when getting with success",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (domain.Truck, error) {
						return domain.Truck{}, nil
					},
				},
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{
							ID:           1,
							TruckID:      1,
							Origin:       "90",
							Destination:  "90",
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
			want: domain.Trip{
				ID:           1,
				TruckID:      1,
				Origin:       "90",
				Destination:  "90",
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
			tr := NewTripService(tt.fields.tripRepository, tt.fields.truckService, tt.fields.TripUpdater)
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
		tripUpdater    ITripUpdater
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
			name: "should return an error when GetTrip fails",
			fields: fields{
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{}, errors.New("generic error")
					},
				},
			},
			args: args{
				location: domain.Location{},
			},
			wantErr: true,
		},
		{
			name: "should return an error when UpsertTrip fails",
			fields: fields{
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{}, nil
					},
					UpsertTripMock: func(trip *domain.Trip) error {
						return errors.New("failed to update")
					},
				},
				tripUpdater: TripUpdaterMock{
					UpdateTripMock: func(currentTrip domain.Trip, location domain.Location) domain.Trip {
						return domain.Trip{}
					},
				},
			},
			args: args{
				location: domain.Location{},
			},
			wantErr: true,
		},
		{
			name: "should nil when upserting with success",
			fields: fields{
				tripRepository: &TripRepositoryMock{
					GetTripMock: func(truckID int) (domain.Trip, error) {
						return domain.Trip{}, nil
					},
					UpsertTripMock: func(trip *domain.Trip) error {
						return nil
					},
				},
				tripUpdater: TripUpdaterMock{
					UpdateTripMock: func(currentTrip domain.Trip, location domain.Location) domain.Trip {
						return domain.Trip{}
					},
				},
			},
			args: args{
				location: domain.Location{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTripService(
				tt.fields.tripRepository,
				tt.fields.truckService,
				tt.fields.tripUpdater,
			)
			if err := tr.UpdateTrip(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("TripService.UpdateTrip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TripRepositoryMock struct {
	GetTripMock    func(truckID int) (domain.Trip, error)
	UpsertTripMock func(trip *domain.Trip) error
}

func (t TripRepositoryMock) GetTrip(truckID int) (domain.Trip, error) {
	return t.GetTripMock(truckID)
}
func (t TripRepositoryMock) UpsertTrip(trip *domain.Trip) error {
	return t.UpsertTripMock(trip)
}

type TripUpdaterMock struct {
	UpdateTripMock func(currentTrip domain.Trip, location domain.Location) domain.Trip
}

func (t TripUpdaterMock) UpdateTrip(currentTrip domain.Trip, location domain.Location) domain.Trip {
	return t.UpdateTripMock(currentTrip, location)
}
