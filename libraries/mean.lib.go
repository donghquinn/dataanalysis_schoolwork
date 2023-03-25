package libraries

import (
	"fmt"
	"strconv"
)

func GetFloatMean(data []string, columnName string) {
	total := 0.0
	max := 0.0
	min := 1.0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Conver error: ", err)
		}

		if max < number {

			max = number
		}

		if min > number {

			min = number
		}

		total += number
	}

	mean := total / float64(len(data))

	fmt.Printf("%s Data. Mean: %f, Max: %f, Min: %f\n ", columnName, mean, max, min)
}

func GetIntMean(data []string, columnName string) int {
	var total int = 0
	max := 0
	min := 0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.Atoi(data[i])
		initmin, minErr := strconv.Atoi(data[1])

		if minErr != nil {
			fmt.Println("Min Convert Error: ", minErr)
		}

		min = initmin

		if err != nil {
			fmt.Println("Conver error: ", err)
		}

		if max < number {
			max = number
		}

		if min > number {
			min = number
		}

		total += number
	}

	mean := (total) / (len(data))

	fmt.Printf("Column: %s, Mean: %d, Max: %d, Min: %d\n", columnName, mean, max, min)

	return mean
}
