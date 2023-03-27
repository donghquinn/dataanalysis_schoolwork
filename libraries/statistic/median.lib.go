package statistic

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

var (
	median1 float64
	median2 float64
)

func GetMedian(data []string, dataName string) (float64, float64) {
	totalDataNumber := len(data) - 1

	fmt.Println("Total Data Number: ", totalDataNumber)

	convertedDataArray := make([]float64, totalDataNumber)

	for i := 1; i < len(data); i += 1 {
		element, elementErr := strconv.ParseFloat(data[i], 64)

		if elementErr != nil {
			fmt.Println("Element Error", elementErr.Error())
		}

		convertedDataArray = append(convertedDataArray, element)
	}

	sort.Float64s(convertedDataArray)

	if totalDataNumber/2 != 0 {
		fmt.Println("Median Index: ", int(math.Ceil(float64(totalDataNumber)/2)))

		median1 = convertedDataArray[int(math.Ceil(float64(totalDataNumber)/2))]

		return median1, math.NaN()
	} else {
		fmt.Println("First Median Index: ", totalDataNumber/2)
		median1 = convertedDataArray[totalDataNumber/2]
		fmt.Println("First Median Number: ", median1)

		fmt.Println("Second Median Index: ", totalDataNumber/2+1)
		median2 = convertedDataArray[totalDataNumber/2+1]
		fmt.Println("Second Median Number: ", median2)

		return median1, median2
	}
}
