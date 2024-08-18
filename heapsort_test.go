package algorithms

import (
	"reflect"
	"testing"
)

func TestHeapSortInt(t *testing.T) {
	type args struct {
		array   []int
		reverse bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1] Happy Case, ascending order",
			args: args{
				array:   []int{5, 6, 6, 4, 3, 10, 1},
				reverse: false,
			},
			want: []int{1, 3, 4, 5, 6, 6, 10},
		},
		{
			name: "[2] Happy Case, descending order",
			args: args{
				array:   []int{5, 6, 6, 4, 3, 10, 1},
				reverse: true,
			},
			want: []int{10, 6, 6, 5, 4, 3, 1},
		},
		{
			name: "[3] Edge Case, empty list",
			args: args{
				array:   []int{},
				reverse: false,
			},
			want: []int{},
		},
		{
			name: "[4] Edge Case, zero values",
			args: args{
				array:   []int{0, 0, 0},
				reverse: false,
			},
			want: []int{0, 0, 0},
		},
		{
			name: "[5] Boundary Case, negative integers in set",
			args: args{
				array:   []int{-5, -6, -6, -4, -3, -10, 1, 10},
				reverse: false,
			},
			want: []int{-10, -6, -6, -5, -4, -3, 1, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSortInt(tt.args.array, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSortInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
