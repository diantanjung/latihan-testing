package dependencyinjection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    DefaultEngine
		want int
	}{
		{
			name: "test default engine",
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DefaultEngine{}
			if got := e.MaxSpeed(); got != tt.want {
				t.Errorf("DefaultEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    TurboEngine
		want int
	}{
		{
			name: "test turbo engine",
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := TurboEngine{}
			if got := e.MaxSpeed(); got != tt.want {
				t.Errorf("TurboEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}



type MockEngine struct{
	mock.Mock
}

func (e MockEngine) MaxSpeed() int {
	args := e.Called()
	return args.Get(0).(int)
}

func TestCar_Speed_With_Mock(t *testing.T) {
	mock := new(MockEngine)
	car := Car{
		Engine: mock,
	}
	mock.On("MaxSpeed").Return(9).Times(1)

	assert.Equal(t,10, car.Speed())
	mock.AssertExpectations(t)
}

type FakeEngine struct{}

func (e FakeEngine) MaxSpeed() int {
	return 9
}

func TestCar_Speed_With_Fake(t *testing.T) {

	type fields struct {
		Engine Engine
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "default engine case",
			fields: fields{Engine: &DefaultEngine{}},
			want: 70,
		},
		{
			name: "turbo engine case",
			fields: fields{Engine: &TurboEngine{}},
			want: 100,
		},
		{
			name: "turbo engine case with fake",
			fields: fields{Engine: &FakeEngine{}},
			want: 10,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Engine: tt.fields.Engine,
			}
			if got := c.Speed(); got != tt.want {
				t.Errorf("Car.Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}
