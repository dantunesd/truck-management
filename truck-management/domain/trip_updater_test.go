package domain

import (
	"reflect"
	"testing"
)

func TestTripUpdater_UpdateTrip(t *testing.T) {
	type args struct {
		trip     Trip
		location Location
	}
	tests := []struct {
		name string
		args args
		want Trip
	}{
		{
			"should return an ONGOING Trip when receiving a empty trip and the EngineState is ON",
			args{
				trip: Trip{
					ID:      0,
					TruckID: 0,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  ON,
					CurrentSpeed: 120,
					Latitude:     92,
					Longitude:    63,
					EngineHours:  2,
					Odometer:     200,
				},
			},
			Trip{
				TruckID:      1,
				Origin:       "92 63",
				Destination:  "92 63",
				State:        ONGOING,
				Odometer:     200,
				EngineHours:  2,
				AverageSpeed: 100,
			},
		},
		{
			"should return a FINISHED Trip when receiving a empty and the EngineState is OFF ",
			args{
				trip: Trip{
					ID:      0,
					TruckID: 0,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  OFF,
					CurrentSpeed: 120,
					Latitude:     92,
					Longitude:    63,
					EngineHours:  2,
					Odometer:     200,
				},
			},
			Trip{
				TruckID:      1,
				Origin:       "92 63",
				Destination:  "92 63",
				State:        FINISHED,
				Odometer:     200,
				EngineHours:  2,
				AverageSpeed: 100,
			},
		},
		{
			"should return an ONGOING Trip when the current trip is finished and the EngineState is ON",
			args{
				trip: Trip{
					ID:           1,
					TruckID:      1,
					Origin:       "92 63",
					Destination:  "92 63",
					State:        FINISHED,
					Odometer:     200,
					EngineHours:  2,
					AverageSpeed: 100,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  ON,
					CurrentSpeed: 120,
					Latitude:     100,
					Longitude:    110,
					EngineHours:  4,
					Odometer:     500,
				},
			},
			Trip{
				ID:           1,
				TruckID:      1,
				Origin:       "100 110",
				Destination:  "100 110",
				State:        ONGOING,
				Odometer:     500,
				EngineHours:  4,
				AverageSpeed: 125,
			},
		},
		{
			"should return a FINISHED Trip when the current trip is finished and the EngineState is OFF ",
			args{
				trip: Trip{
					ID:           1,
					TruckID:      1,
					Origin:       "92 63",
					Destination:  "92 63",
					State:        FINISHED,
					Odometer:     200,
					EngineHours:  2,
					AverageSpeed: 100,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  OFF,
					CurrentSpeed: 120,
					Latitude:     100,
					Longitude:    110,
					EngineHours:  4,
					Odometer:     500,
				},
			},
			Trip{
				ID:           1,
				TruckID:      1,
				Origin:       "100 110",
				Destination:  "100 110",
				State:        FINISHED,
				Odometer:     500,
				EngineHours:  4,
				AverageSpeed: 125,
			},
		},
		{
			"should return an ONGOING Trip when the current trip is ongoing and the EngineState is ON",
			args{
				trip: Trip{
					ID:           1,
					TruckID:      1,
					Origin:       "92 63",
					Destination:  "92 63",
					State:        ONGOING,
					Odometer:     200,
					EngineHours:  2,
					AverageSpeed: 100,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  ON,
					CurrentSpeed: 120,
					Latitude:     100,
					Longitude:    110,
					EngineHours:  2,
					Odometer:     200,
				},
			},
			Trip{
				ID:           1,
				TruckID:      1,
				Origin:       "92 63",
				Destination:  "100 110",
				State:        ONGOING,
				Odometer:     400,
				EngineHours:  4,
				AverageSpeed: 100,
			},
		},
		{
			"should return a FINISHED Trip when the current trip is ongoing and the EngineState is OFF ",
			args{
				trip: Trip{
					ID:           1,
					TruckID:      1,
					Origin:       "92 63",
					Destination:  "92 63",
					State:        ONGOING,
					Odometer:     200,
					EngineHours:  2,
					AverageSpeed: 100,
				},
				location: Location{
					ID:           1,
					TruckID:      1,
					EngineState:  OFF,
					CurrentSpeed: 120,
					Latitude:     100,
					Longitude:    110,
					EngineHours:  2,
					Odometer:     200,
				},
			},
			Trip{
				ID:           1,
				TruckID:      1,
				Origin:       "92 63",
				Destination:  "100 110",
				State:        FINISHED,
				Odometer:     400,
				EngineHours:  4,
				AverageSpeed: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTripUpdater()
			if got := tr.UpdateTrip(tt.args.trip, tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TripUpdater.UpdateTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
