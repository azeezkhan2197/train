package join

import (
	"fmt"

	"github.com/azeezkhan2197/train/constant"
	"github.com/azeezkhan2197/train/model"
)

func AppendArray(trainA, trainB []model.BOGIE) []model.BOGIE {
	trainAB := []model.BOGIE{}
	indexA, indexB := 0, 0
	for {
		if indexA >= len(trainA) || indexB >= len(trainB) {
			break
		}
		if trainA[indexA].Distance > trainB[indexB].Distance {
			trainAB = append(trainAB, trainA[indexA])
			indexA++
		} else {
			trainAB = append(trainAB, trainB[indexB])
			indexB++
		}
	}

	if indexA == len(trainA) {
		trainAB = append(trainAB, trainB[indexB:]...)
	} else {
		trainAB = append(trainAB, trainA[indexA:]...)
	}
	return trainAB
}

func ArrivalTrainOrderHyd(bogeyCount int) []model.BOGIE {
	train := []model.BOGIE{}
	RouteA := map[string]int{
		"CHN": 0,
		"SLM": 350,
		"BLR": 550,
		"KRN": 900,
		"HYB": 1200,
		"NGP": 1600,
		"ITJ": 1900,
		"BPL": 2000,
		"AGA": 2500,
		"NDL": 2700,
	}

	RouteB := map[string]int{
		"TVC": 0,
		"SRR": 300,
		"MAQ": 600,
		"MAO": 1000,
		"PNE": 1400,
		"HYB": 2000,
		"NGP": 2400,
		"ITJ": 2700,
		"BPL": 2800,
		"PTA": 3800,
		"NJP": 4200,
		"GHY": 4700,
	}

	for index := 0; index < bogeyCount; index++ {
		var stationCode string
		fmt.Printf("enter the staion code of bogie %d ", index+1)
		Bogie := model.BOGIE{}
		fmt.Scan(&stationCode)
		if RouteA[stationCode] != 0 {
			Distance := RouteA[stationCode] - RouteA[constant.HYDERABADKEY]
			if Distance >= 0 {
				Bogie.Station = stationCode
				Bogie.Distance = Distance
				train = append(train, Bogie)
			}
		} else {
			Distance := RouteB[stationCode] - RouteB[constant.HYDERABADKEY]
			if Distance >= 0 {
				Bogie.Station = stationCode
				Bogie.Distance = Distance
				train = append(train, Bogie)
			}
		}
	}
	return train

}
