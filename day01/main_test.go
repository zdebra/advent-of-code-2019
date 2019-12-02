package main

import "testing"

func TestModule_FuelRequiredToLaunch(t *testing.T) {
	type fields struct {
		Mass float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "example 1",
			fields: fields{
				Mass: 12,
			},
			want:   2,
		},
		{
			name:   "example 2",
			fields: fields{
				Mass: 14,
			},
			want:   2,
		},
		{
			name:   "example 3",
			fields: fields{
				Mass: 1969,
			},
			want:   966,
		},
		{
			name:   "example 4",
			fields: fields{
				Mass: 100756,
			},
			want:   50346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				Mass: tt.fields.Mass,
			}
			if got := m.FuelRequiredToLaunch(); got != tt.want {
				t.Errorf("FuelRequiredToLaunch() = %v, want %v", got, tt.want)
			}
		})
	}
}
