package libraries

import "fmt"

func RemodeledCount(data []string) (int, int, int) {
	recentCount := 0
	oldCount := 0
	noneCount := 0

	for i := 0; i < len(data); i += 1 {
		if data[i] == "Recent" {
			recentCount += 1
		} else if data[i] == "Old" {
			oldCount += 1
		} else {
			noneCount += 1
		}
	}

	fmt.Printf("Remodeling Data - Total Data: %d Recent: %d, Old: %d, None: %d\n", len(data), recentCount, oldCount, noneCount)

	return recentCount, oldCount, noneCount
}

func ChasCount(data []string) (int, int) {
	countZero := 0
	countOne := 0

	for i := 0; i < len(data); i += 1 {
		if data[i] == "0" {
			countZero += 1
		} else if data[i] == "1" {
			countOne += 1
		} else {
			fmt.Println("Unexpected Data: ", data[i])
		}
	}

	fmt.Printf("CHAS Data - Total Data: %d '1': %d, '0': %d\n", len(data), countOne, countZero)

	return countOne, countZero
}
