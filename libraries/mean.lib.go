package libraries

import (
	"fmt"
	"strconv"
)

func GetFloatMean(data []string, columnName string) {
	total := 0.0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Conver error: ", err)
		}

		total += number
	}

	mean := total / float64(len(data))

	fmt.Printf("%s Data. Mean: %f\n ", columnName, mean)
}

func GetIntMean(data []string, columnName string) int {
	var total int = 0

	for i := 1; i < len(data); i += 1 {
		number, err := strconv.Atoi(data[i])

		if err != nil {
			fmt.Println("Conver error: ", err)
		}

		total += number
	}

	mean := (total) / (len(data))

	fmt.Printf("Column: %s, Mean: %d\n", columnName, mean)

	return mean
}
