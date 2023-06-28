package join

import (
	"reflect"
	"testing"

	"github.com/azeezkhan2197/train/model"
)

func TestAppendArray(t *testing.T) {
	type args struct {
		trainA []model.BOGIE
		trainB []model.BOGIE
	}
	tests := []struct {
		name string
		args args
		want []model.BOGIE
	}{
		{
			name: "simple case",
			args: args{
				trainA: []model.BOGIE{{"HYB", 1}, {"NGP", 3}},
				trainB: []model.BOGIE{{"NDL", 2}},
			},
			want: []model.BOGIE{{"NDL", 2}, {"HYB", 1}, {"NGP", 3}},
		},
		{
			name: "b is bigger than a ",
			args: args{
				trainA: []model.BOGIE{{"HYB", 1}},
				trainB: []model.BOGIE{{"NDL", 2}, {"NGP", 3}},
			},
			want: []model.BOGIE{{"NDL", 2}, {"NGP", 3}, {"HYB", 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendArray(tt.args.trainA, tt.args.trainB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrivalTrainOrderHyd(t *testing.T) {
	type args struct {
		bogeyCount int
	}
	tests := []struct {
		name string
		args args
		want []model.BOGIE
	}{
		{
			name: "success",
			args: args{
				bogeyCount: 0,
			},
			want: []model.BOGIE{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrivalTrainOrderHyd(tt.args.bogeyCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrivalTrainOrderHyd() = %v, want %v", got, tt.want)
			}
		})
	}
}
