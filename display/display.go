package display

import (
	"fmt"

	"github.com/azeezkhan2197/train/constant"
	"github.com/azeezkhan2197/train/model"
)

func PrintArray(train []model.BOGIE) []model.BOGIE {
	indexA := 0
	for indexA < len(train) {
		fmt.Printf(" %s ", train[indexA].Station)
		if train[indexA].Station == constant.HYDERABADKEY {
			train = append(train[:indexA], train[indexA+1:]...)
			continue
		}
		indexA++
	}
	fmt.Println()
	return train
}
