package main

import (
	"fmt"

	"github.com/azeezkhan2197/train/display"
	"github.com/azeezkhan2197/train/join"
	"github.com/azeezkhan2197/train/sort"
)

func main() {

	var trainACount, trainBCount int
	fmt.Println("\nenter the number of bogies for TRAIN A")
	fmt.Scan(&trainACount)
	trainA := join.ArrivalTrainOrderHyd(trainACount)
	fmt.Println("\nenter the number of bogies for TRAIN B")
	fmt.Scan(&trainBCount)
	trainB := join.ArrivalTrainOrderHyd(trainBCount)

	fmt.Print("\nARRIVAL   TRAIN_A ENGINE ")
	trainA = display.PrintArray(trainA)
	
	fmt.Print("\nARRIVAL   TRAIN_B ENGINE ")
	trainB = display.PrintArray(trainB)
	
	sort.Quick_sort(trainA, 0, len(trainA)-1)
	sort.Quick_sort(trainB, 0, len(trainB)-1)

	trainAB := join.AppendArray(trainA, trainB)
	fmt.Print("\nDEPARTURE   TRAIN_AB ENGINE ENGINE ")
	for _, val := range trainAB {
		fmt.Printf(" %s ", val.Station)
	}
}
