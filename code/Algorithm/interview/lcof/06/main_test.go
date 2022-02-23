// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了96.36%的用户
// 通过测试用例：24 / 24

package main

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
