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
		medival1 = data[totalDataNumber/2]

		return medival1, ""
	} else {
		fmt.Println("First Medival Index: ", int(math.Ceil(float64(totalDataNumber)/2)))
		medival1 = data[int(math.Ceil(float64(totalDataNumber)/2))]
		fmt.Println("First Medival Number: ", medival1)

		fmt.Println("Second Medival Index: ", int(math.Ceil(float64(totalDataNumber)/2))+1)

		medival2 = data[int(math.Ceil(float64(totalDataNumber)/2))+1]
		fmt.Println("Second Medival Number: ", medival2)

		return medival1, medival2
	}
}
