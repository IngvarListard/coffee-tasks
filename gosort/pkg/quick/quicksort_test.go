package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				list: []int{119, 281, 762, 21, 1, 99, 0, 615, 2275, 8, 4, 2},
			},
			want: []int{0, 1, 2, 4, 8, 21, 99, 119, 281, 615, 762, 2275},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Sort(tt.args.list); !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", tt.args.list, tt.want)
			}
		})
	}
}
