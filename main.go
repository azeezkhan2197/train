package main

import (
	"fmt"

	"github.com/azeezkhan2197/train/display"
	"github.com/azeezkhan2197/train/join"
	"github.com/azeezkhan2197/train/sort"
)

func main() {

	var tACount, tBCount int
	fmt.Println("enter the number of bogies for TRAIN A")
	fmt.Scan(&tACount)
	trainA := join.ArrivalTrainHyd(tACount)
	fmt.Println("enter the number of bogies for TRAIN B")
	fmt.Scan(&tBCount)
	trainB := join.ArrivalTrainHyd(tBCount)

	fmt.Println()
	fmt.Print("ARRIVAL   TRAIN_A ENGINE ")
	trainA = display.PrintArray(trainA)
	fmt.Println()

	fmt.Print("ARRIVAL   TRAIN_B ENGINE ")
	trainB = display.PrintArray(trainB)
	fmt.Println()

	sort.Quick_sort(trainA, 0, len(trainA)-1)
	sort.Quick_sort(trainB, 0, len(trainB)-1)

	trainAB := join.Append(trainA, trainB)
	fmt.Print("DEPARTURE   TRAIN_AB ENGINE ENGINE ")
	for _, val := range trainAB {
		fmt.Printf(" %s ", val.Station)
	}

	fmt.Println()
}
