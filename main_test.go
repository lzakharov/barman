package main

import "testing"

func TestPlot(t *testing.T) {
	type args struct {
		xs []float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: struct{ xs []float64 }{xs: []float64{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			}},
			want: "▁▁▁▂▂▃▃▄▄▅▅▆▆▇▇█",
		},
		{
			name: "monotone",
			args: struct{ xs []float64 }{xs: []float64{
				1, 1, 1, 1, 1, 1, 1, 1,
			}},
			want: "▅▅▅▅▅▅▅▅",
		},
		{
			name: "floats",
			args: struct{ xs []float64 }{xs: []float64{
				0.66, 0.23, 0.78, 0.72, 0.81, 0.54, 0.76, 0.43, 0.34, 0.66,
			}},
			want: "▆▁▇▆█▄▇▃▂▆",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plot(tt.args.xs); got != tt.want {
				t.Errorf("Plot() = %v, want %v", got, tt.want)
			}
		})
	}
}
