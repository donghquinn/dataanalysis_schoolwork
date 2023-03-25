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

	fmt.Printf("Remodeling Data. Total Data: %d Recent: %d, Old: %d, None: %d\n", len(data), recentCount, oldCount, noneCount)

	return recentCount, oldCount, noneCount
}
