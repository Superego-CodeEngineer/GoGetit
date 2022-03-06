package dynamic

import "testing"

func TestDynamic(t *testing.T) {
	type args struct {
		prices  []int
		weights []int
		left    int
	}
	tests := []struct {
		name      string
		args      args
		wantValue int
	}{
		{
			name: "0-1",
			args: args{
				prices:  []int{10, 8, 20, 5},
				weights: []int{2, 5, 9, 3},
				left:    10,
			},
			wantValue: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := Dynamic(tt.args.prices, tt.args.weights, tt.args.left); gotValue != tt.wantValue {
				t.Errorf("Dynamic() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
