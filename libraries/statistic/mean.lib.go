package statistic

import (
	"fmt"
	"strconv"
)

func GetFloatMean(data []string, columnName string) {
	total := 0.0
	max := 0.0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Conver error: ", err.Error())
		}

		if max < number {
			max = number
		}

		total += number
	}

	mean := total / float64(len(data))

	medival1, medival2 := GetMedian(data, columnName)

	if medival2 == "" {
		fmt.Printf("Column: %s, Average: %f, Max: %f, Medival: %s \n", columnName, mean, max, medival1)
	} else {
		fmt.Printf("Column: %s, Average: %f, Max: %f, Medival1: %s, Medival2: %s \n", columnName, mean, max, medival1, medival2)
	}

	// fmt.Printf("%s Data. Average: %f, Max: %f\n ", columnName, mean, max)
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
	median1, median2 := GetMedian(data, columnName)

	if median2 == "" {
		fmt.Printf("Column: %s, Average: %d, Max: %d, Median: %s \n", columnName, mean, max, median1)
	} else {
		fmt.Printf("Column: %s, Average: %d, Max: %d, Median1: %s, Median2: %s \n", columnName, mean, max, median1, median2)
	}
}
