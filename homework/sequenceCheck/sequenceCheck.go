package sequenceCheck

import "sort"

func Solution(inputArray []int) int {
	sort.Ints(inputArray)

	for index, value := range inputArray {
		if index+1 != value {
			return 0
		}
	}

	return 1
}
