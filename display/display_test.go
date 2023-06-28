package display

import (
	"reflect"
	"testing"

	"github.com/azeezkhan2197/train/model"
)

func TestPrintArray(t *testing.T) {
	type args struct {
		train []model.BOGIE
	}
	tests := []struct {
		name string
		args args
		want []model.BOGIE
	}{
		{
			name: "success",
			args: args{
				train: []model.BOGIE{{"HYB",1},{"NDL",5},{"KRN",10}},
			},
			want: []model.BOGIE{{"NDL",5},{"KRN",10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintArray(tt.args.train); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
