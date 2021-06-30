package domain

import "testing"

func TestIsValidTruck(t *testing.T) {
	type args struct {
		truck         Truck
		existingTruck Truck
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return error when exists a truck with the same lisence plate",
			args: args{
				truck: Truck{
					LicensePlate: "ABC1234",
				},
				existingTruck: Truck{
					LicensePlate: "ABC1234",
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when exists a truck with the same eld id",
			args: args{
				truck: Truck{
					EldID:        "00001234",
					LicensePlate: "NEW1234",
				},
				existingTruck: Truck{
					EldID:        "00001234",
					LicensePlate: "OLD1234",
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when trucks are empty",
			args: args{
				truck:         Truck{},
				existingTruck: Truck{},
			},
			wantErr: true,
		},
		{
			name: "should return nil when it's a new truck",
			args: args{
				truck: Truck{
					EldID:        "00001234",
					LicensePlate: "NEW1234",
				},
				existingTruck: Truck{
					EldID:        "",
					LicensePlate: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsValidTruck(tt.args.truck, tt.args.existingTruck); (err != nil) != tt.wantErr {
				t.Errorf("IsValidTruck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
