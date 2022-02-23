package insertsort

import (
	"reflect"
	"testing"
)

func Test_insertsort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[3, 1, 5, 4, 2]",
			args: args{
				[]int{3, 1, 5, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertsort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
