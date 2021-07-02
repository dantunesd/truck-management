package domain

import "testing"

func TestIsValidTruck(t *testing.T) {
	type fields struct {
		LicensePlateChecker LicensePlateChecker
		EldChecker          EldChecker
	}
	type args struct {
		truck Truck
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return error when exists a truck with the same lisence plate",
			fields: fields{
				LicensePlateChecker: &LicensePlateCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return true
					},
				},
				EldChecker: &EldCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return false
					},
				},
			},
			args: args{
				truck: Truck{
					LicensePlate: "ABC1234",
					EldID:        "00001234",
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when exists a truck with the same eld id",
			fields: fields{
				LicensePlateChecker: &LicensePlateCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return false
					},
				},
				EldChecker: &EldCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return true
					},
				},
			},
			args: args{
				truck: Truck{
					LicensePlate: "NEW1234",
					EldID:        "00001234",
				},
			},
			wantErr: true,
		},
		{
			name: "should return nil when it's a new truck",
			fields: fields{
				LicensePlateChecker: &LicensePlateCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return false
					},
				},
				EldChecker: &EldCheckerMock{
					IsAlreadyInUseMock: func(string) bool {
						return false
					},
				},
			},
			args: args{
				truck: Truck{
					EldID:        "00001234",
					LicensePlate: "NEW1234",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TruckValidator{
				licensePlateChecker: tt.fields.LicensePlateChecker,
				eldChecker:          tt.fields.EldChecker,
			}
			if err := v.IsValidTruck(tt.args.truck); (err != nil) != tt.wantErr {
				t.Errorf("IsValidTruck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type LicensePlateCheckerMock struct {
	IsAlreadyInUseMock func(string) bool
}

func (l *LicensePlateCheckerMock) IsAlreadyInUse(licensePlate string) bool {
	return l.IsAlreadyInUseMock(licensePlate)
}

type EldCheckerMock struct {
	IsAlreadyInUseMock func(string) bool
}

func (l *EldCheckerMock) IsAlreadyInUse(eldID string) bool {
	return l.IsAlreadyInUseMock(eldID)
}
