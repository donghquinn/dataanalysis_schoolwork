package statistic

import (
	"fmt"
	"strconv"
)

func GetFloatMean(data []string, columnName string) {
	total := 0.0
	max := 0.0
	// min := 0.0

	// tmpArray := make([]float64, 2)

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.ParseFloat(data[i], 64)

		// tmpArray = append(tmpArray, number)

		if err != nil {
			fmt.Println("Convert error: ", err.Error())
		}

		if max < number {
			max = number
		}

		// if i >= 2 {
		// 	fmt.Printf("tmpArray: %v, number: %f\n", tmpArray, number)
		// 	decideMin := math.Min(tmpArray[0], number)

		// 	min = decideMin
		// }

		total += number
	}

	mean := total / float64(len(data))

	median := GetMedian(data, columnName)

	fmt.Printf("Column: %s, Average: %f, Max: %f, Median: %f \n", columnName, mean, max, median)
}

func GetIntMean(data []string, columnName string) {
	var total int = 0
	max := 0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.Atoi(data[i])

		if err != nil {
			fmt.Println("String to Int Convert error: ", err.Error())
		}

		if max < number {
			max = number
		}

		total += number
	}

	mean := (total) / (len(data))
	median := GetMedian(data, columnName)

	fmt.Printf("Column: %s, Average: %d, Max: %d, Median: %f \n", columnName, mean, max, median)
}
