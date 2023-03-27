package statistic

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

var (
	medival1 string
	medival2 string
)

func GetMedival(data []string, dataName string) {
	totalDataNumber := len(data) - 1

	fmt.Println("Total Data Number: ", totalDataNumber)

	convertedDataArray := make([]float64, totalDataNumber)

	for i := range data {
		element, elementErr := strconv.ParseFloat(data[i], 64)

		if elementErr != nil {
			fmt.Println("Element Error", elementErr.Error())
		}

		convertedDataArray = append(convertedDataArray, element)

		sort.Float64s(convertedDataArray)
	}

	if totalDataNumber/2 != 0 {
		medival1 = data[totalDataNumber/2]
	} else {
		fmt.Println("First Medival Index: ", int(math.Ceil(float64(totalDataNumber)/2)))
		medival1 = data[int(math.Ceil(float64(totalDataNumber)/2))]
		fmt.Println("First Medival Number: ", medival1)

		fmt.Println("Second Medival Index: ", int(math.Ceil(float64(totalDataNumber)/2))+1)

		medival2 = data[int(math.Ceil(float64(totalDataNumber)/2))+1]
		fmt.Println("Second Medival Number: ", medival2)
	}
}
