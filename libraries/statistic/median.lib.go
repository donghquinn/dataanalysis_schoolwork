package statistic

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

var (
	median1 string
	median2 string
)

func GetMedian(data []string, dataName string) (string, string) {
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
	// fmt.Printf("Converted: %v", convertedDataArray)

	if totalDataNumber/2 != 0 {
		fmt.Println("Median Index: ", int(math.Ceil(float64(totalDataNumber)/2)))

		median1 = data[int(math.Ceil(float64(totalDataNumber)/2))]

		return median1, ""
	} else {
		fmt.Println("First Median Index: ", totalDataNumber/2)
		median1 = data[totalDataNumber/2]
		fmt.Println("First Median Number: ", median1)

		fmt.Println("Second Median Index: ", totalDataNumber/2+1)
		median2 = data[totalDataNumber/2+1]
		fmt.Println("Second Median Number: ", median2)

		return median1, median2
	}
}
