package sort

import (
	"testing"

	"github.com/azeezkhan2197/train/model"
)

func TestPartition(t *testing.T) {
	type args struct {
		a    []model.BOGIE
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				a:    []model.BOGIE{{"HYD", 1}, {"HYD", 1}},
				low:  0,
				high: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Partition(tt.args.a, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuick_sort(t *testing.T) {
	type args struct {
		a    []model.BOGIE
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				a:    []model.BOGIE{{"HYB",1},{"KRN",10}},
				low:  0,
				high: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quick_sort(tt.args.a, tt.args.low, tt.args.high)
		})
	}
}
