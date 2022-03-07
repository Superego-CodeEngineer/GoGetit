package backtracking

import "testing"

func TestBacktracking(t *testing.T) {
	type args struct {
		prices  []int
		weights []int
		item    int
		left    int
		value   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0-1",
			args: args{
				prices:  []int{10, 8, 20, 5},
				weights: []int{2, 5, 9, 3},
				item:    0,
				left:    10,
				value:   0,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Backtracking(tt.args.prices, tt.args.weights, tt.args.item, tt.args.left, tt.args.value); got != tt.want {
				t.Errorf("Backtracking() = %v, want %v", got, tt.want)
			}
		})
	}
}
