package basic

import "testing"

func TestAbsolute(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name :	"negative case",
			args :	args{-5},
			want :	5,
		},
		{
			name :	"positive case",
			args :	args{5},
			want :	5,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Absolute(tt.args.num); got != tt.want {
				t.Errorf("Absolute() = %v, want %v", got, tt.want)
			}
		})
	}
}
