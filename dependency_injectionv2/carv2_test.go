package dependency_injectionv2

import (
	"testing"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    *DefaultEngine
		want int
	}{
		// TODO: Add test cases.
		{
			name: "default engine should have 50 maxspeed",
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.MaxSpeed(); got != tt.want {
				t.Errorf("DefaultEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    *TurboEngine
		want int
	}{
		// TODO: Add test cases.
		{
			name: "turbo engine should have 100 maxspeed",
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.MaxSpeed(); got != tt.want {
				t.Errorf("TurboEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO:Make fake Test
type FakeEngine struct{}

func (e FakeEngine) MaxSpeed() int {
	return 5
}

func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "default engine should have 50 maxspeed",
			fields: fields{Speeder: &DefaultEngine{}},
			want:   50,
		},
		{
			name:   "turbo engine should have 100 maxspeed",
			fields: fields{Speeder: &TurboEngine{}},
			want:   100,
		},
		{
			name:   "speed should be 20 when use maxspeed is less than 10",
			fields: fields{Speeder: &FakeEngine{}},
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Speeder: tt.fields.Speeder,
			}
			if got := c.Speed(); got != tt.want {
				t.Errorf("Car.Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}
