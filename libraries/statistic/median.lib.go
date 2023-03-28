package statistic

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	median float64
)

// 중앙값
func GetMedian(data []string, dataName string) float64 {
	// column 열은 빼고
	totalDataNumber := len(data)

	// fmt.Println("Total Data Number: ", totalDataNumber)

	convertedDataArray := make([]float64, totalDataNumber)

	for i := 1; i < len(convertedDataArray); i += 1 {
		element, elementErr := strconv.ParseFloat(data[i], 64)

		if elementErr != nil {
			fmt.Println("Element Error", elementErr.Error())
		}

		convertedDataArray[i] = element
	}

	sort.Float64s(convertedDataArray)

	middleNumber := len(convertedDataArray) / 2

	if len(convertedDataArray)%2 != 0 {
		fmt.Println("Median Index: ", middleNumber)

		median = convertedDataArray[middleNumber]

		return median
	} else {
		fmt.Println("First Median Index: ", middleNumber)
		first := convertedDataArray[middleNumber]
		fmt.Println("First Median Number: ", first)

		fmt.Println("Second Median Index: ", (middleNumber)+1)
		second := convertedDataArray[(middleNumber)+1]
		fmt.Println("Second Median Number: ", second)

		median = (first + second) / 2

		return median
	}
}
