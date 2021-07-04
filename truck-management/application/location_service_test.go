package application

import (
	"errors"
	"reflect"
	"testing"
	"truck-management/truck-management/domain"
)

func TestLocationService_CreateLocation(t *testing.T) {
	type fields struct {
		locationRepository ILocationRepository
		truckService       ITruckService
	}
	type args struct {
		truckID  int
		location domain.Location
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Location
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
				locationRepository: &LocationRepositoryMock{
					CreateLocationMock: func(truckID int, location *domain.Location) error {
						return nil
					},
				},
			},
			args: args{
				location: domain.Location{},
			},
			want:    domain.Location{},
			wantErr: true,
		},
		{
			name: "should return an error when CreateLocation returns any error",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (*domain.Truck, error) {
						return &domain.Truck{}, nil
					},
				},
				locationRepository: &LocationRepositoryMock{
					CreateLocationMock: func(truckID int, location *domain.Location) error {
						return errors.New("failed to create truck")
					},
				},
			},
			args: args{
				location: domain.Location{},
			},
			want:    domain.Location{},
			wantErr: true,
		},
		{
			name: "should return LocationEntity when creating with success",
			fields: fields{
				truckService: TruckServiceMock{
					GetTruckMock: func(ID int) (*domain.Truck, error) {
						return &domain.Truck{}, nil
					},
				},
				locationRepository: &LocationRepositoryMock{
					CreateLocationMock: func(truckID int, location *domain.Location) error {
						location.ID = 1
						location.TruckID = truckID
						location.CreatedAt = timeNow
						return nil
					},
				},
			},
			args: args{
				truckID: 5,
				location: domain.Location{
					EldID:        "id",
					EngineState:  "ON",
					CurrentSpeed: 100,
					Latitude:     1000,
					Longitude:    1000,
					EngineHours:  1,
					Odometer:     100,
				},
			},
			want: domain.Location{
				ID:           1,
				TruckID:      5,
				EldID:        "id",
				EngineState:  "ON",
				CurrentSpeed: 100,
				Latitude:     1000,
				Longitude:    1000,
				EngineHours:  1,
				Odometer:     100,
				CreatedAt:    timeNow,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLocationService(tt.fields.locationRepository, tt.fields.truckService)
			got, err := l.CreateLocation(tt.args.truckID, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocationService.CreateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationService.CreateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

type LocationRepositoryMock struct {
	CreateLocationMock  func(truckID int, location *domain.Location) error
	GetLastLocationMock func(truckID int) (*domain.Location, error)
}

func (l LocationRepositoryMock) CreateLocation(truckID int, location *domain.Location) error {
	return l.CreateLocationMock(truckID, location)
}
func (l LocationRepositoryMock) GetLastLocation(truckID int) (*domain.Location, error) {
	return l.GetLastLocationMock(truckID)
}

type TruckServiceMock struct {
	GetTruckMock func(ID int) (*domain.Truck, error)
}

func (t TruckServiceMock) GetTruck(ID int) (*domain.Truck, error) {
	return t.GetTruckMock(ID)
}
