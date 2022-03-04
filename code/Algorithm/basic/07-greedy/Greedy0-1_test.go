package greedy

import "testing"

func Test_greedy(t *testing.T) {
	type args struct {
		s   []int
		p   []int
		bag int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "0-1",
			args: args{
				s:   []int{10, 2, 20, 5},
				p:   []int{2, 5, 4, 8},
				bag: 10,
			},
			wantAns: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := greedy(tt.args.s, tt.args.p, tt.args.bag); gotAns != tt.wantAns {
				t.Errorf("greedy() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
